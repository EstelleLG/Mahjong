// An implementation of Conway's Game of Life in Commandline
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
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
	f.board[y][x] = b
}

func (f *Field) isAlive(x, y int) bool {
	x = (x + f.width) % f.width
	y = (y + f.height) % f.height
	return f.board[y][x]
}

func (f *Field) next(x, y int) bool {
	alive := 0
	for i := -1; i < 1; i++ {
		for j := -1; j < 1; j++ {
			if (j != 0 || i != 0) && f.isAlive(x+i, y+j) {
				alive++
			}
		}
	}
	return alive == 3 || (alive == 2 && f.isAlive(x, y))
}

type Life struct {
	old, new      *Field
	width, height int
}

func (life *Life) Step() {
	for x := 0; x < life.width; x++ {
		for y := 0; y < life.height; y++ {
			life.new.Set(x, y, life.old.next(x, y))
		}
	}
	life.old, life.new = life.new, life.old
}

func (life *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < life.height; y++ {
		for x := 0; x < life.width; x++ {
			b := byte(' ')
			if life.old.isAlive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	for i := 0; i < 100; i++ {
		buf.WriteByte('-')
	}
	return buf.String()
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
	l := randomNewLife(40, 15)
	for i := 0; i < 3; i++ {
		l.Step()
		fmt.Print("\x0c", l)
		time.Sleep(time.Second / 30)
	}
}
