package board

type Knight struct {
	core *piece
}

func NewKnight(side Color) Piece {
	c := &piece{
		name: "Knight",
		side: side,
	}
	return &Knight{
		core: c,
	}
}

func (p *Knight) Name() string {
	return p.core.name
}

func (p *Knight) Ascii() string {
	if p.core.side == White {
		return string('n')
	} else {
		return string('N')
	}
}

func (p *Knight) Unicode() string {
	return string(FillKnight)
}

func (p *Knight) Side() Color {
	return p.core.side
}

func (p *Knight) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, *loc)
	motions := [8][2]int{{-2, 1}, {-2, -1}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {1, 2}, {-1, 2}}
	for _, motion := range motions {
		targetx := loc.x + motion[0]
		targety := loc.y + motion[1]

		if targetx < 0 || targetx >= boardSize {
			continue
		}
		if targety < 0 || targety >= boardSize {
			continue
		}

		target := board.squares[targetx][targety]

		if target.Occupied() {
			if target.Occupant().Side() == p.Side() {
				continue
			}
		}

		moves = append(moves, target)
	}
	return moves
}

//NOOP
func (p *Knight) Move() {}
