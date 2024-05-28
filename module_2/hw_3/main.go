package main

import "fmt"

func main() {
	var size, price int
	var contains bool
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&price)
	for _, val := range arr {
		if val == price {
			contains = true
			break
		}
	}
	fmt.Println(contains)
}
