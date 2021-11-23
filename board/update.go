package board

import (
	"fmt"
	"sort"

	"github.com/gdamore/tcell/v2"
)

type Input int

const (
	Left Input = iota
	Up
	Right
	Down
	In
	Out
	Noop
)

type SortDirection int

type SortAxis int

const (
	Rank SortAxis = iota
	File
)

const (
	Forward SortDirection = iota
	Reverse
)

func (b *Board) Update(ev *tcell.EventKey) {
	input := inputFromKeypress(ev)
	if b.activeCursor().mode == Insert {
		b.applyInsert(input)
	} else {
		b.applySelect(input)
	}
}

func (b *Board) applyInsert(i Input) {
	switch i {
	case Left, Right, Up, Down:
		b.moveCursor(i)
	case In:
		//move
	case Out:
		b.activeCursor().switchMode()
	}
}

func (b *Board) applySelect(i Input) {
	switch i {
	case Left, Right, Up, Down:
		b.moveCursor(i)
	case In:
		b.activeCursor().switchMode()
	case Out:
		break
	}
}

func (b *Board) moveCursor(i Input) {
	var axis SortAxis
	var dir SortDirection
	if i == Left || i == Right {
		axis = File
	} else {
		axis = Rank
	}
	if i == Left || i == Up {
		dir = Reverse
	} else {
		dir = Forward
	}
	sortedSqs := sortSquares(b.Moves(), axis, dir)
	fmt.Println(sortedSqs)
	loc := b.activeCursor().loc
	var target Square
	for i, sq := range sortedSqs {
		if sq == loc {
			if i+1 != len(sortedSqs) {
				target = sortedSqs[i+1]
			} else {
				target = sortedSqs[0]
			}
		}
	}
	b.activeCursor().loc = b.squares[target.x][target.y]
}

func sortSquares(sqs []Square, axis SortAxis, dir SortDirection) []Square {
	sqsPrime := make([]Square, len(sqs))
	fmt.Println(sqs)
	_ = copy(sqsPrime, sqs)
	if axis == Rank {
		sortRankFirst(sqsPrime)
	} else {
		sortFileFirst(sqsPrime)
	}
	if dir == Reverse {
		for i, j := 0, len(sqsPrime)-1; i < j; i, j = i+1, j-1 {
			sqsPrime[i], sqsPrime[j] = sqsPrime[j], sqsPrime[i]
		}
	}
	return sqsPrime
}

func sortFileFirst(sqs []Square) {
	sort.Slice(sqs, func(i, j int) bool {
		if sqs[i].y < sqs[j].y {
			return true
		} else if sqs[i].y > sqs[j].y {
			return false
		}
		return sqs[i].x < sqs[j].x
	})
}

func sortRankFirst(sqs []Square) {
	sort.Slice(sqs, func(i, j int) bool {
		if sqs[i].x < sqs[j].x {
			return true
		} else if sqs[i].x > sqs[j].x {
			return false
		}
		return sqs[i].y < sqs[j].y
	})
}

func inputFromKeypress(ev *tcell.EventKey) Input {
	action := Noop
	switch ev.Key() {
	case tcell.KeyRune:
		switch ev.Rune() {
		case 'j', 'J', 'a', 'A':
			action = Left
		case 'k', 'K', 'w', 'W':
			action = Up
		case 'l', 'L', 'd', 'D':
			action = Right
		case ';', ':', 's', 'S':
			action = Down
		case 'q', 'Q':
			action = Out
		case 'e', 'E':
			action = In
		case ' ':
			action = Out
		}
	case tcell.KeyLeft:
		action = Left
	case tcell.KeyUp:
		action = Up
	case tcell.KeyRight:
		action = Right
	case tcell.KeyDown:
		action = Down
	case tcell.KeyEnter:
		action = In
	}
	return action
}

/*
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func write(str string) {
	log := "/home/imw/src/rtc/log.txt"
	var f *os.File
	f, err := os.OpenFile(log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = fmt.Fprintf(w, "%s\n", str)
	check(err)
	w.Flush()
}
*/
