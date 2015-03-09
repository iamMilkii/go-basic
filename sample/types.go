package main

import "fmt"

func main() {
	fmt.Println("#Number")
	fmt.Println("2 + 2 =",2 + 2)
	fmt.Println("3 + 3 =",3.0 + 3.0)
	fmt.Println("25.30 - 22.12 = ",25.30 - 22.12)

	fmt.Println("#String")
	fmt.Println("Hello " + "world")
	fmt.Println(len("Hello world"))
	fmt.Println("Hello world"[4])
	fmt.Println(string("Hello"[0]))

	fmt.Println("#Booleans")
	fmt.Println("true && true = ", true && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false)
}
