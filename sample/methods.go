package main

import "fmt"

type Rect struct {
	width, height int
}

func (r *Rect) area() int {
	return r.width * r.height
}

func (r *Rect) perim() int {
	return (2 * r.width) + (2 * r.height)
}

func main() {
	r := Rect{width : 20, height : 30}
	
	fmt.Println("area :", r.area())
	fmt.Println("perim :", r.perim())
}
