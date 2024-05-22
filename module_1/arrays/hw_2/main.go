package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Переменные для размера массива и строки ввода
	var size int
	var line string
	var newSlice []int

	// Чтение размера массива
	fmt.Scan(&size)

	// Чтение второй строки, содержащей элементы массива
	line, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	line = strings.TrimSpace(line) // Удаление пробельных символов в начале и конце строки

	// Разделение строки на отдельные элементы
	parts := strings.Split(line, " ")

	// Создание массива и преобразование строк в числа
	arr := make([]int, size)
	for i, part := range parts {
		arr[i], _ = strconv.Atoi(part) // Преобразование строки в число и добавление в массив
	}

	// Формирование нового слайса, сначала добавляем все ненулевые элементы
	for _, val := range arr {
		if val != 0 {
			newSlice = append(newSlice, val)
		}
	}

	// Затем добавляем все нулевые элементы
	for _, val := range arr {
		if val == 0 {
			newSlice = append(newSlice, val)
		}
	}

	// Вывод элементов нового слайса через пробел
	for _, val := range newSlice {
		fmt.Print(val, " ")
	}
}
