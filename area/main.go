package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triagle struct {
	base   float64
	height float64
}

func main() {
	s := square{sideLength: 5}
	t := triagle{base: 4, height: 6}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triagle) getArea() float64 {
	return 0.5 * t.base * t.height
}
