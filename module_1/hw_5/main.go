package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Book представляет книгу с ISBN, названием и годом издания
type Book struct {
	ISBN  int
	Title string
	Year  int
}

func main() {
	var size int
	fmt.Scan(&size)
	books := make([]Book, 0, size)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < size; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		// Извлечение ISBN
		spaceIndex := strings.Index(line, " ")
		isbnStr := line[:spaceIndex]
		isbn, _ := strconv.Atoi(isbnStr)

		// Извлечение года издания
		lastSpaceIndex := strings.LastIndex(line, " ")
		yearStr := line[lastSpaceIndex+1:]
		year, _ := strconv.Atoi(yearStr)

		// Извлечение названия книги
		title := line[spaceIndex+1 : lastSpaceIndex]

		// Создание новой книги и добавление в массив
		book := Book{ISBN: isbn, Title: title, Year: year}
		books = append(books, book)
	}

	// Сортировка книг по году и названию
	sort.SliceStable(books, func(i, j int) bool {
		if books[i].Year == books[j].Year {
			return books[i].Title < books[j].Title
		}
		return books[i].Year < books[j].Year
	})

	// Вывод отсортированных книг
	for _, book := range books {
		fmt.Printf("%d %s %d\n", book.ISBN, book.Title, book.Year)
	}
}
