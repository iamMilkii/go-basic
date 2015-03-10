package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	
	s := Person{"Bobo", 21}	
	fmt.Println(s)

	y := Person{name : "Bebe", age :20}
	fmt.Println(y)

	var z Person
	z.name = "kobeap"
	z.age = 24
	fmt.Println("My name is", z.name)
	fmt.Println("Age", z.age)

	sp := &z
	sp.age = 30
	fmt.Println(sp.age)
	fmt.Println(z.age)
}
