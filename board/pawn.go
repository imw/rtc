package board

type Pawn struct {
	core  *piece
	moved bool
}

func NewPawn(side Color) Piece {
	c := &piece{
		name: "Pawn",
		side: side,
	}
	return &Pawn{
		core: c,
	}
}

func (p *Pawn) Name() string {
	return p.core.name
}

func (p *Pawn) Ascii() string {
	if p.core.side == White {
		return string('P')
	} else {
		return string('p')
	}
}

func (p *Pawn) Unicode() string {
	return string(FillPawn)
}

func (p *Pawn) Side() Color {
	return p.core.side
}

func (p *Pawn) GetMoves(board Board, loc Square) []Square {
	moves := []Square{}
	x := loc.x
	y := loc.y
	//move two forward from start position
	if p.moved == false {
		if board.squares[x][y-1].Unoccupied() && board.squares[x][y-2].Unoccupied() {
			moves = append(moves, board.squares[x][y-2])
		}
	}

	//move one forward
	if board.squares[x][y-1].Unoccupied() {
		moves = append(moves, board.squares[x][y-1])
	}

	//take left
	if board.squares[x-1][y-1].Occupied() && (board.squares[x-1][y-1].Occupant().Side() != p.Side()) {
		moves = append(moves, board.squares[x-1][y-1])
	}

	//take right
	if board.squares[x+1][y-1].Occupied() && (board.squares[x+1][y-1].Occupant().Side() != p.Side()) {
		moves = append(moves, board.squares[x+1][y-1])
	}

	return moves

}

func (p *Pawn) Move() {
	p.moved = true
}
