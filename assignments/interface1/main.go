package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func printArea(sh shape) {
	area := sh.getArea()
	fmt.Println(area)
}

func main() {
	sq1 := square{
		sideLength: 4.0,
	}

	tr1 := triangle{
		height: 1.0,
		base:   5.0,
	}

	printArea(sq1)
	printArea(tr1)
}
