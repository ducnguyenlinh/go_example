package main

import "fmt"

func one_dimensional_slice() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("update:", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// slice operator
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
}

func multi_dimensional_slice() {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func main() {
	one_dimensional_slice()

	multi_dimensional_slice()
}
