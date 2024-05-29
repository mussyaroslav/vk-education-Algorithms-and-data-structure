package main

import "fmt"

func main() {
	var size, maxValue, maxKey int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	myMap := make(map[int]int)
	for _, v := range arr {
		myMap[v]++
	}
	for k, v := range myMap {
		if v > maxValue {
			maxValue = v
			maxKey = k
		}
	}
	if maxValue > len(arr)/2 {
		fmt.Println(maxKey)
	} else {
		fmt.Println(-1)
	}
}
