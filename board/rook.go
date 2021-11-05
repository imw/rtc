package board

type Rook struct {
	core *piece
}

func NewRook(side Color) Piece {
	c := &piece{
		name: "Rook",
		side: side,
	}
	return &Rook{
		core: c,
	}
}

func (p *Rook) Name() string {
	return p.core.name
}

func (p *Rook) Ascii() string {
	if p.core.side == White {
		return string('r')
	} else {
		return string('R')
	}
}

func (p *Rook) Unicode() string {
	return string(FillRook)
}

func (p *Rook) Side() Color {
	return p.core.side
}
