package board

type Queen struct {
	core *piece
}

func NewQueen(side Color) Piece {
	c := &piece{
		name: "Queen",
		side: side,
	}
	return &Queen{
		core: c,
	}
}

func (p *Queen) Name() string {
	return p.core.name
}

func (p *Queen) Ascii() string {
	if p.core.side == White {
		return string('q')
	} else {
		return string('Q')
	}
}

func (p *Queen) Unicode() string {
	return string(FillQueen)
}

func (p *Queen) Side() Color {
	return p.core.side
}
