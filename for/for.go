package main

import "fmt"

func loop1(min int, max int) int {
	i := min
	for i <= max {
		fmt.Println(i)
		i = i + 1
	}

	return i
}

func loop2(min int, max int) {
	for j:= min; j<= max; j++ {
		fmt.Println(j)
	}
}

func loop3(min int, max int) {
	for n := min; n <= max; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func main() {
	fmt.Println(loop1(1, 3))

	loop2(7, 9)

	for {
		fmt.Println("loop")
		break
	}

	loop3(0, 5)
}
