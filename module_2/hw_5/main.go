package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	myMap := map[string]int{}
	maxValue := 0
	for _, v := range strings.Split(string(line), "") {
		myMap[v]++
	}
	for _, v := range myMap {
		if v > maxValue {
			maxValue = v
		}
	}
	fmt.Println(maxValue)
}
