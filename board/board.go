package board

import (
	"fmt"
	"strconv"
)

const boardSize = 8

type Board struct {
	squares [boardSize][boardSize]Square
}

func New() *Board {
	b := new(Board)
	b.squares = [boardSize][boardSize]Square{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			name := string('A'+i) + strconv.Itoa(j+1)
			s := Square{
				x:        i,
				y:        j,
				name:     name,
				occupant: nil,
			}
			b.squares[i][j] = s
		}
	}
	return b
}

func (*Board) Print() {
	fmt.Println(WhiteBishop)
}
