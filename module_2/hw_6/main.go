package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	trueFalse := false
	arr := make([]string, 2)
	for i := 0; i < 2; i++ {
		fmt.Scan(&arr[i])
	}
	for i, val := range arr {
		val = strings.ReplaceAll(val, "\"", "")
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		arr[i] = string(runes)
		if arr[0] == arr[1] {
			trueFalse = true
		}
	}
	fmt.Println(trueFalse)
}
