package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение первой строки и преобразование в количество элементов
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)

	// Чтение второй строки и преобразование в массив чисел
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	arr := make([]int, n)
	for i, part := range parts {
		arr[i], _ = strconv.Atoi(part)
	}

	// Чтение третьей строки и преобразование в элемент
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	element, _ := strconv.Atoi(line)

	// Подсчет количества элементов, не равных заданному числу
	result := countNonMatchingElements(arr, element)

	// Вывод результата
	fmt.Println(result)
}

// countNonMatchingElements принимает массив целых чисел и число element.
// Возвращает количество элементов, не равных element.
func countNonMatchingElements(arr []int, element int) int {
	count := 0
	for _, value := range arr {
		if value != element {
			count++
		}
	}
	return count
}
