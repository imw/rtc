package board

//Knight implements the Piece interface for knights
type Knight struct {
	core *piece
}

//NewKnight returns a new Knight
func NewKnight(side Color) Piece {
	c := &piece{
		name: "Knight",
		side: side,
	}
	return &Knight{
		core: c,
	}
}

//Name returns the name of this piece
func (p *Knight) Name() string {
	return p.core.name
}

//ASCII returns the Ascii representation for this piece
func (p *Knight) ASCII() string {
	if p.core.side == White {
		return string('n')
	}
	return string('N')
}

//Unicode returns the unicode rune for this piece
func (p *Knight) Unicode() string {
	return string(FillKnight)
}

//Side return's this piece's side
func (p *Knight) Side() Color {
	return p.core.side
}

//ValidMoves returns a slice of squares representing valid moves for this
//piece, given board and location
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

//Move is a noop for Knights
func (p *Knight) Move() {}
