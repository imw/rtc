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

	return c.loc.Occupant().GetMoves(c.board, c.loc)
}

// record moves for playback/rollback
