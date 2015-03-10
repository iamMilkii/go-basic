package main

import "fmt"

func main() {
	m := make(map[string]int)
	
	m["a1"] = 11
	m["a2"] = 12
	
	fmt.Println("map :", m)
	
	val1 := m["a1"]
	
	fmt.Println("value1 :", val1)
	fmt.Println("len :", len(m))
	
	delete(m, "a2")
	
	fmt.Println("map :", m)

	_, prs := m["a2"]
	
	fmt.Println("prs :", prs)

	n := map[int]string{1 : "ap", 2 : "qqw", 3 : "kkk"}

	fmt.Println("map :", n)
}
