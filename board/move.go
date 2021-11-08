package board

type CursorMode int

const (
	Select CursorMode = iota
	Insert
)

type Cursor struct {
	board Board
	loc   Square
	mode  CursorMode
	color Color
}

//returns slice of legal squares - all occupied squares in select mode, all
//legal moves in insert mode
func (c *Cursor) Choices() []Square {
	if c.mode == Select {
		sqs := []Square{}
		for _, rank := range c.board.squares {
			for _, sq := range rank {
				if sq.Occupied() && sq.Occupant().Side() == c.color {
					sqs = append(sqs, sq)
				}
			}
		}
		return sqs
	}

	return c.loc.Occupant().ValidMoves(c.board, c.loc)
}

func search(xmotion binop, ymotion binop, board Board, loc Square) []Square {
	moves := []Square{}
	x := loc.x
	y := loc.y
	p := loc.occupant
	for i := 1; i < boardSize; i++ {
		targetx := xmotion(x, i)
		targety := ymotion(y, i)
		if targetx > boardSize || targety > boardSize {
			break
		}
		target := board.squares[xmotion(x, i)][ymotion(y, i)]
		if target.Occupied() {
			if target.Occupant().Side() != p.Side() {
				moves = append(moves, target)
			}
			break
		} else {
			moves = append(moves, target)
		}
	}
	return moves
}

func seekForward(board Board, loc Square) []Square {
	return search(id, sub, board, loc)
}

func seekReverse(board Board, loc Square) []Square {
	return search(id, add, board, loc)
}

func seekLeft(board Board, loc Square) []Square {
	return search(sub, id, board, loc)
}

func seekRight(board Board, loc Square) []Square {
	return search(add, id, board, loc)
}

func seekForwardL(board Board, loc Square) []Square {
	return search(sub, add, board, loc)
}

func seekForwardR(board Board, loc Square) []Square {
	return search(add, add, board, loc)
}

func seekReverseL(board Board, loc Square) []Square {
	return search(sub, sub, board, loc)
}

func seekReverseR(board Board, loc Square) []Square {
	return search(add, sub, board, loc)
}

type binop func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func id(a, b int) int {
	return a
}

// TODO: record moves for playback/rollback
