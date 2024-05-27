package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	newArr := make([]int, 0)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]%2 == 0 {
			newArr = append(newArr, arr[i])
		}
	}
	if len(newArr) > 0 {
		fmt.Println(newArr[0])
	} else {
		fmt.Println(-1)
	}
}
