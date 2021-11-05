package board

type Pawn struct {
	core *piece
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
