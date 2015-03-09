package main

import "fmt"

func main() {
	var name string = "Apichat Eakwongsa."
	nickName := "Kobeap."
	fmt.Println("My name is", name)
	fmt.Println("My nick name", nickName)

	const age int = 24
	fmt.Printf("Age %d. \n", age)
	
	var input int
	fmt.Print("Enter number : ")
	fmt.Scanf("%d", &input)
	output := input * 5
	fmt.Printf("%d * 5 = %d \n", input, output)
}
