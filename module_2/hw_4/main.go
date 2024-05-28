package main

import (
	"fmt"
	"sort"
)

func main() {
	var size, num int
	fmt.Scan(&size)
	arr := make([]int, size)
	newArr := make([]int, 0)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&num)
	newArr = append(newArr, arr...)
	newArr = append(newArr, num)
	sort.Ints(newArr)
	for i, _ := range newArr {
		if newArr[i] == num {
			fmt.Print(i)
			break
		}
	}
}
