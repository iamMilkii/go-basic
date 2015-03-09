package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[2] = 30
	fmt.Println(x)

	y := [3]float64{2, 3, 4}
	var count float64
	for i, value := range y {
		count = count + value
		if i < (len(y) - 1) {
			fmt.Print(i)
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println(count)

	z := [5]int {
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(z)
}
