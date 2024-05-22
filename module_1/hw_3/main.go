package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
}
