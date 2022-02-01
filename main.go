package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/kelseyhightower/envconfig"

	"0x539.lol/rtc/board"
	"0x539.lol/rtc/display"
	"0x539.lol/rtc/game"
	"0x539.lol/rtc/util"
)

type Config struct {
	HostA string `default:"localhost"`
	PortA string `default:"1234"`
	HostB string `default:"localhost"`
	PortB string `default:"4321"`
	ID    string `default:"A"`
}

func main() {
	var c Config
	err := envconfig.Process("rtc", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	util.Write(fmt.Sprintf("Config: %v", c))

	var b *board.Board

	var remotehost string
	var remoteport string
	var localhost string
	var localport string
	if c.ID == "A" {
		localhost = c.HostA
		localport = c.PortA
		remotehost = c.HostB
		remoteport = c.PortB
		b = board.New(board.White)
	} else {
		localhost = c.HostB
		localport = c.PortB
		remotehost = c.HostA
		remoteport = c.PortA
		b = board.New(board.Black)
	}

	g, gr := game.New(b)

	//register RPC
	rpc.Register(gr)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", localhost+":"+localport)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	defer l.Close()
	util.Write(fmt.Sprintf("Listening with listener: %v", l))

	//TODO loading/waiting screen

	//setup display
	encoding.Register()
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	display.Loading(s)

	dialCount := 1
	for {
		//init client for peer
		util.Write(fmt.Sprintf("Dial #%d", dialCount))
		client, err := rpc.DialHTTP("tcp", remotehost+":"+remoteport)
		if err != nil {
			util.Write(fmt.Sprintf("error attempting to dial %s:%s: %s", remotehost, remoteport, err))
			time.Sleep(time.Duration(dialCount) * time.Second)
		} else {
			g.SetClient(client)
			break
		}
		dialCount = dialCount + 1
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			}
		}
	}

	display.Greeting(s)
	go func(b *board.Board, s tcell.Screen) {
		t := time.NewTicker(10 * time.Millisecond)
		for range t.C {
			display.Render(b, s)
		}
	}(b, s)

	//TODO something in board needs to send from channel
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			} else {
				move := b.ProcessEvent(ev)
				var exit int
				if move.Seq > 0 {
					g.SendMove(move)
					gr.DoMove(move, &exit)
				}

				display.Render(b, s)
			}
		}
	}
}
