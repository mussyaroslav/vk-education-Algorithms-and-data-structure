package main

import (
	"fmt"
)

func main() {
	var size int
	var evenIndex int
	fmt.Scan(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	// Если элемент четный, меняем его местами с элементом на позиции evenIndex.
	// Это перемещает четное число в начало массива.
	for i := 0; i < size; i++ {
		if arr[i]%2 == 0 {
			arr[evenIndex], arr[i] = arr[i], arr[evenIndex]
			evenIndex++
		}
	}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
}
