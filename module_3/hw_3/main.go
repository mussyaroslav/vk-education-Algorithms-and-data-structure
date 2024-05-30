package main

import (
	"fmt"
)

func main() {
	var size, k int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&k)

	maxMultiply := 0

	for i := 0; i <= size-k; i++ {
		num := 1
		for j := i; j < i+k; j++ {
			num *= arr[j]
		}
		if num > maxMultiply {
			maxMultiply = num
		}
	}

	fmt.Println(maxMultiply)
}
