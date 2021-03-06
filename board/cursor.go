package board

//CursorMode enumerates whether cursor is set to select or insert
type CursorMode int

//Select chooses a piece to move, Insert chooses a target
const (
	Select CursorMode = iota
	Insert
)

//Cursor encapsulates information about move selection
type Cursor struct {
	loc    *Square
	target *Square
	mode   CursorMode
	color  Color
}

func (c *Cursor) switchMode() {
	if c.mode == Select {
		c.mode = Insert
		c.target = c.loc
	} else {
		c.mode = Select
		c.target = nil
	}
}

//returns slice of legal squares - all occupied squares in select mode, all
//legal moves in insert mode
func (c *Cursor) choices(board Board) []Square {
	if c.mode == Select {
		sqs := []Square{}
		for _, rank := range board.squares {
			for _, sq := range rank {
				if sq.Occupied() && sq.Occupant().Side() == c.color {
					sqs = append(sqs, sq)
				}
			}
		}
		return sqs
	}
	return c.loc.Occupant().ValidMoves(board, c.loc)
}

func search(xmotion binop, ymotion binop, board Board, loc *Square, limit int) []Square {
	moves := []Square{}
	moves = append(moves, *loc)
	x := loc.x
	y := loc.y
	p := loc.occupant
	for i := 1; i <= limit; i++ {
		targetx := xmotion(x, i)
		targety := ymotion(y, i)
		if targetx >= boardSize || targetx < 0 {
			break
		}
		if targety >= boardSize || targety < 0 {
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

func seekForward(board Board, loc *Square, limit int) []Square {
	return search(id, sub, board, loc, limit)
}

func seekReverse(board Board, loc *Square, limit int) []Square {
	return search(id, add, board, loc, limit)
}

func seekLeft(board Board, loc *Square, limit int) []Square {
	return search(sub, id, board, loc, limit)
}

func seekRight(board Board, loc *Square, limit int) []Square {
	return search(add, id, board, loc, limit)
}

func seekForwardL(board Board, loc *Square, limit int) []Square {
	return search(sub, add, board, loc, limit)
}

func seekForwardR(board Board, loc *Square, limit int) []Square {
	return search(add, add, board, loc, limit)
}

func seekReverseL(board Board, loc *Square, limit int) []Square {
	return search(sub, sub, board, loc, limit)
}

func seekReverseR(board Board, loc *Square, limit int) []Square {
	return search(add, sub, board, loc, limit)
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
