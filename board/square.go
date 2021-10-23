package board

type Square struct {
	x        int
	y        int
	name     string
	occupant Piece
	color    Color
}

func (s *Square) Indices() (int, int) {
	return s.x, s.y
}

func (s *Square) Occupant() Piece {
	return s.occupant
}

func (s *Square) Color() Color {
	return s.color
}
