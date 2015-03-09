package main

import "fmt"

func main() {
	var input int
	fmt.Println("# For Loop")
	fmt.Print("Enter number : ")
	fmt.Scanf("%d", &input)
	for i := 1 ; i <= input	; i++ {
		fmt.Println("one", i)
	}

	fmt.Println("=====")	

	i := 1
	for i <= 10 {
		fmt.Println("two", i)
		i++
	}

	fmt.Println("# if else")
	fmt.Print("Enter number :")
	fmt.Scanf("%d", &input)
	if input % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	i = input % 2
	switch i {
		case 0 : fmt.Println("haha")
		default : fmt.Println("lol")
	}
}
