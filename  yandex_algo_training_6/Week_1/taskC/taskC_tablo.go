package main

import (
	"fmt"
	"go/scanner"
	"strings"
)

// func compress(array [][]string) {
// 	result := array[0]
// 	for _, str := range array[1:] {
// 		if str !=
// 	}
// }

func main() {
	var n int
	var line string

	fmt.Scan(&n)

	array := make([][]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&line)

		elements := strings.Split(line, "")

		array[i] = elements
	}

	for _, v := range array {
		fmt.Println(v)
	}
	for _, row := range array {
		for _, v := range row {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
