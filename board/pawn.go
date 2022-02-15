package board

//Pawn implements the Piece interface for Pawns
type Pawn struct {
	core  *piece
	moved bool
}

//NewPawn returns a new Pawn
func NewPawn(side Color) Piece {
	c := &piece{
		name: "Pawn",
		side: side,
	}
	return &Pawn{
		core: c,
	}
}

//Name returns the name of this piece
func (p *Pawn) Name() string {
	return p.core.name
}

//ASCII returns the Ascii representation for this piece
func (p *Pawn) ASCII() string {
	if p.core.side == White {
		return string('P')
	}
	return string('p')
}

//Unicode returns the unicode rune for this piece
func (p *Pawn) Unicode() string {
	return string(FillPawn)
}

//Side return's this piece's side
func (p *Pawn) Side() Color {
	return p.core.side
}

//ValidMoves returns a slice of squares representing valid moves for this
//piece, given board and location
func (p *Pawn) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, *loc)
	if p.core.side == White {
		if p.moved == false {
			moves = append(moves, seekForward(board, loc, 2)...)
		}
		moves = append(moves, seekForward(board, loc, 1)...)
	} else {
		if p.moved == false {
			moves = append(moves, seekReverse(board, loc, 2)...)
		}
		moves = append(moves, seekReverse(board, loc, 1)...)
	}

	for i, v := range moves {
		if p.opponentOccupied(v) {
			moves = removeSquare(moves, i)
		}
	}

	moves = append(moves, p.attackMoves(board, loc)...)
	return moves

}

func (p *Pawn) attackMoves(board Board, loc *Square) []Square {
	var yop binop
	if p.core.side == White {
		yop = sub
	} else {
		yop = add
	}
	moves := []Square{}
	for _, v := range search(sub, yop, board, loc, 1) {
		if p.opponentOccupied(v) {
			moves = append(moves, v)
		}
	}
	for _, v := range search(add, yop, board, loc, 1) {
		if p.opponentOccupied(v) {
			moves = append(moves, v)
		}
	}

	return moves
}

func (p *Pawn) opponentOccupied(s Square) bool {
	if s.Occupied() {
		if s.Occupant().Side() != p.core.side {
			return true
		}
	}
	return false
}

func removeSquare(sqs []Square, index int) []Square {
	return append(sqs[:index], sqs[index+1:]...)
}

//Move returns whether this piece has moved at least once
func (p *Pawn) Move() {
	p.moved = true
}
