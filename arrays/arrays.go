package main

import "fmt"

func one_dimensionaly_array() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get a[4]:", a[4])

	fmt.Println("len:", len(a))
}

func multi_dimensional_array() {
	b := [5]int{1, 3, 5, 7, 9}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i:= 0; i<2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func main() {
	one_dimensionaly_array()

	multi_dimensional_array()
}
