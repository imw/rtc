package board

import "github.com/gdamore/tcell/v2"

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

func (b *Board) Update(ev tcell.EventKey) {
	input := inputFromKeypress(ev)
	if b.activeCursor().mode == Insert {
		b.applyInsert(input)
	} else {
		b.applySelect(input)
	}
}

func (b *Board) applyInsert(i Input) {
	c := b.activeCursor()
	//if input is motion, seek within choices (with wraparound)
	//if input is in, move
	//if input is out, return to select mode

}

func (b *Board) applySelect(i Input) {
	//if input is motion, seek within choices (with wraparound)
	//if input is in, enter insert mode
	//if input is out, do nothing
}

func inputFromKeypress(ev tcell.EventKey) Input {
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
