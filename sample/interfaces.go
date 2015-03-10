package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Square struct {
	width, height float64
}

func (s Square) area() float64 {
	return s.width * s.height
}

func (s Square) perim() float64 {
	return (2 * s.width) + (2 * s.height)
}

type Circle struct {
	redius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("area :", g.area())
	fmt.Println("perim :", g.perim())
}

func main() {
	s := Square{width : 8, height :5}
	c := Circle{redius : 4}
	measure(s)
	measure(c)
}
