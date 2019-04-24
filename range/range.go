package main

import (
	"fmt"
	"reflect"
)

func sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func range_maps(kvs map[string]string) {
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println("i===>", reflect.TypeOf(i))
		fmt.Println("c===>", reflect.TypeOf(c))
		fmt.Println(i, c)
	}
}
func main() {
	nums := []int{2, 3, 4}
	fmt.Println("sum:", sum(nums))

	for i, num := range nums{
		if num == 2 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	range_maps(kvs)
}
