// An implementation of Conway's Game of Life in Commandline
package main

import "fmt"

type Field struct {
	s    [][]bool
	w, h int
}

func newField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

func main() {
	fmt.Print("HERE")
}
