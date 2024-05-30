package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	maxNum := 0
	minNum := arr[0]
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	fmt.Println(maxNum * minNum)
}
