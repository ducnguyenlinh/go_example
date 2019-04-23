package main

import "fmt"

func main() {
	// init map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// delete
	delete(m, "k2")
	fmt.Println("map:", m)

	// declare and initialize
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)
}
