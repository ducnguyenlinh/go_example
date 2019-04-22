package main

import "fmt"

func check_even_odd(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}

func check_divisible(a int, b int) bool {
	if a%b == 0 {
		return true
	}

	return false
}

func check_multiple_digits(num int) {
	if num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}

func main() {
	n := 7
	if check_even_odd(n) {
		fmt.Println(n, " is even")
	} else {
		fmt.Println(n, " is odd")
	}

	a := 8
	b := 4
	if check_divisible(a, b) {
		fmt.Println(a, " is divisible by ", b)
	}

	num := -2
	check_multiple_digits(num)
}
