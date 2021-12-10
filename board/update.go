package board

import (
	"fmt"
	"sort"

	"0x539.lol/rtc/util"
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
	util.Write(fmt.Sprintf("Active Cursor: %v", b.activeCursor()))
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
		b.moveReticle(i)
	case In:
		b.move(i)
		b.activeCursor().switchMode()
		b.switchCursor()
		util.Write(fmt.Sprintf("board: %v", b))
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

func (b *Board) move(i Input) {
	util.Write("move\n")
	loc := b.activeCursor().loc
	tgt := b.activeCursor().target
	util.Write(fmt.Sprintf("moving %v to %v", loc, tgt))
	p := loc.occupant
	util.Write(fmt.Sprintf("address of p: %v", &p))
	loc.occupant = nil
	tgt.occupant = p
	p.Move()
	util.Write(fmt.Sprintf("after move: %v to %v", loc, tgt))
}

func (b *Board) moveReticle(i Input) {
	util.Write("move reticle\n")
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
	util.Write(fmt.Sprintf("Sorted Squares: %v", sortedSqs))
	tgt := b.activeCursor().target
	util.Write(fmt.Sprintf("Target: %v", tgt))
	var target *Square
	for i, sq := range sortedSqs {
		if sq == *tgt {
			if i+1 != len(sortedSqs) {
				target = &sortedSqs[i+1]
			} else {
				target = &sortedSqs[0]
			}
		}
	}
	util.Write(fmt.Sprintf("moving %v to %v", tgt, target))
	b.activeCursor().target = &b.squares[target.x][target.y]
}

func (b *Board) moveCursor(i Input) {
	util.Write("move cursor\n")
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
	util.Write(fmt.Sprintf("Sorted squares: %v", sortedSqs))
	loc := b.activeCursor().loc
	var target *Square
	for i, sq := range sortedSqs {
		if sq == *loc {
			if i+1 != len(sortedSqs) {
				target = &sortedSqs[i+1]
			} else {
				target = &sortedSqs[0]
			}
		}
	}
	util.Write(fmt.Sprintf("moving %v to %v", loc, target))
	b.activeCursor().loc = &b.squares[target.x][target.y]
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
