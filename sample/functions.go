package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func a() (int, int) {
	return 2, 3
}

func plusAll(args ...int) (int, int) {
	var sum int
	for _, value := range args {
		sum += value	
	}
	return sum, len(args)
} 

func main() {
	fmt.Println("3 + 4 =", plus(3, 4))
	fmt.Println("1 + 2 + 3 =", plusPlus(1, 2, 3))
	
	va1, va2 := a()
	fmt.Printf("multiple return : %d and %d \n", va1, va2)
	
	vs1, vs2 := plusAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("plus %d value = %d \n", vs2, vs1)

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println("add 10, 20 : ",add(10, 20))
}
