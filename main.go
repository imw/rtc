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

//Config encapsulates runtime config of RTC application for use by envconfig
type Config struct {
	HostA string `default:"localhost"`
	PortA string `default:"1234"`
	HostB string `default:"localhost"`
	PortB string `default:"4321"`
	ID    string `default:"A"`
	Debug bool   `default:"false"`
}

type GameEngine struct {
	ID       string
	game     *game.Game
	rpc      *game.GameRPC
	listener net.Listener
}

type ExitCode int

const (
	Success ExitCode = iota
	Failure
)

func main() {
	var c Config
	err := envconfig.Process("rtc", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	os.Exit(int(entrypoint(c)))
}

func setup(c Config) (GameEngine, error) {

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

	mux := http.NewServeMux()
	http.DefaultServeMux = mux
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", localhost+":"+localport)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	util.Write(fmt.Sprintf("Listening with listener: %v", l))

	dialCount := 1
	for {
		//init client for peer
		util.Write(fmt.Sprintf("Dial #%d", dialCount))
		client, e := rpc.DialHTTP("tcp", remotehost+":"+remoteport)
		if e != nil {
			util.Write(fmt.Sprintf("error attempting to dial %s:%s: %s", remotehost, remoteport, e))
			if dialCount > 10 {
				return GameEngine{}, e
			}
		} else {
			g.SetClient(client)
			break
		}
		dialCount = dialCount + 1
		/*
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					s.Fini()
					return Success
				}
			}
		*/
		time.Sleep(time.Duration(1) * time.Second)
	}

	return GameEngine{c.ID, g, gr, l}, nil

}

func entrypoint(c Config) ExitCode {

	util.Debug = c.Debug
	util.Write(fmt.Sprintf("Config: %v", c))

	//setup display
	encoding.Register()
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		return Failure
	}

	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		return Failure
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	display.Loading(s)

	engine, e := setup(c)
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		return Failure
	}
	defer engine.listener.Close()

	go func(g *game.Game, s tcell.Screen) {
		t := time.NewTicker(10 * time.Millisecond)
		for range t.C {
			display.Render(g, s)
		}
	}(engine.game, s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				return Failure
			} else {
				util.Write(fmt.Sprintf("Event registered: %v", ev))
				move := engine.game.RecieveEvent(ev)
				util.Write(fmt.Sprintf("Move generated: %v", move))
				var rpcExitCode int
				if move.Seq > 0 {
					engine.game.SendMove(move)
					engine.rpc.DoMove(move, &rpcExitCode)
				}

				display.Render(engine.game, s)
			}
		}
	}
}
