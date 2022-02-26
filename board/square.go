package board

//Square enapsulates information about a square on a board, including
//a reference to any occupant
type Square struct {
	x        int
	y        int
	name     string
	occupant Piece
	color    Color
}

//Indices returns the array indices of a square
func (s *Square) Indices() (int, int) {
	return s.x, s.y
}

//Name returns the name of a square in standard chess annotation
func (s *Square) Name() string {
	return s.name
}

//SetOccupant sets a square to contain a reference to a piece
func (s *Square) SetOccupant(p *Piece) {
	if p == nil {
		s.occupant = nil
	} else {
		s.occupant = *p
	}
}

//Occupant returns an occupying piece, if any
func (s *Square) Occupant() Piece {
	return s.occupant
}

//Occupied returns whether a square is occupied
func (s *Square) Occupied() bool {
	return s.occupant != nil
}

/*
func (s *Square) Unoccupied() bool {
	return s.occupant != nil
}
*/

//Color returns the background color of a square
func (s *Square) Color() Color {
	return s.color
}
