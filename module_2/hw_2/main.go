package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	stack := []rune{}

	for _, char := range s {
		// Если стек не пуст и верхний элемент совпадает с текущим символом
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1] // Удаляем верхний элемент
		} else {
			stack = append(stack, char) // Добавляем текущий символ в стек
		}
	}
	fmt.Println(string(stack))
}
