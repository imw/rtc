package board

type Bishop struct {
	core *piece
}

func NewBishop(side Color) Piece {
	c := &piece{
		name: "Bishop",
		side: side,
	}
	return &Bishop{
		core: c,
	}
}

func (p *Bishop) Name() string {
	return p.core.name
}

func (p *Bishop) Ascii() string {
	if p.core.side == White {
		return string('b')
	} else {
		return string('B')
	}
}

func (p *Bishop) Unicode() string {
	return string(FillBishop)
}

func (p *Bishop) Side() Color {
	return p.core.side
}

//TODO: Implement
func (p *Bishop) GetMoves(board Board, loc Square) []Square {
	moves := []Square{}
	return moves
}

//NOOP
func (p *Bishop) Move() {
}
