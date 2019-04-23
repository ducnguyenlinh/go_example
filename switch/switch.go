package main

import (
	"fmt"
	"strconv"
	"time"
)

func switch1(i int) string {
	switch i {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	}
	
	return strconv.Itoa(i)
}

func check_weekday() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

func check_hour() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

func check_type(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm a int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

func main() {
	fmt.Println(switch1(2))

	check_hour()

	check_weekday()

	check_type(true)
	check_type(1)
	check_type("hey")
}
