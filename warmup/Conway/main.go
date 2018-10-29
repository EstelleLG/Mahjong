// An implementation of Conway's Game of Life in Commandline
package main

import (
	"fmt"
	"math/rand"
)

type Field struct {
	board         [][]bool
	width, height int
}

func newField(width, height int) *Field {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return &Field{board: board, width: width, height: height}
}

func (f *Field) Set(x, y int, b bool) {
	f.board[x][y] = b
}

type Life struct {
	old, new      *Field
	width, height int
}

func randomNewLife(width, height int) *Life {
	old := newField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		old.Set(rand.Intn(width), rand.Intn(height), true)
	}
	return &Life{old: old, new: newField(width, height), width: width, height: height}
}

// func newLifeFromConfig(width, height int) *Life {}

func main() {
	fmt.Print("HERE")
}
