package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	
	s := person{"Bobo", 21}	
	fmt.Println(s)

	y := person{name : "Bebe", age :20}
	fmt.Println(y)

	var z person
	z.name = "kobeap"
	z.age = 24
	fmt.Println("My name is", z.name)
	fmt.Println("Age", z.age)

	sp := &z
	sp.age = 30
	fmt.Println(sp.age)
	fmt.Println(z.age)
}
