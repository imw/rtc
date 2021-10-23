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

func (p *Pawn) Ascii() rune {
	if p.core.side == White {
		return 'P'
	} else {
		return 'p'
	}
}

func (p *Pawn) Unicode() rune {
	if p.core.side == White {
		return WhitePawn
	} else {
		return BlackPawn
	}

}
