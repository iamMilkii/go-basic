package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("example :",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set :", s)
	fmt.Println("get :", s[1])
	fmt.Println("len :", len(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println("append :", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy :", c)
	
	l := s[2:5]
	fmt.Println("sl1 :", l)

	l = s[:4]
	fmt.Println("sl2 :", l)

	l = s[3:]
	fmt.Println("sl3 :", l)
	
	t := []string{"h", "i", "j", "k"}
	fmt.Println("dcl :", t)

	sl2d := make([][]string, 3)
	fmt.Println("slice2d :", sl2d)

	sl2d[0] = make([]string, 3)
	sl2d[0][0] = "q"
	sl2d[0][1] = "s"
	sl2d[0][2] = "f"
	
	sl2d[1] = make([]string, 2)
	sl2d[1][0] = "m"
	sl2d[1][1] = "n"

	sl2d[2] = make([]string, 1)
	sl2d[2][0] = "r"

	fmt.Println("slice2d set :", sl2d)
}
