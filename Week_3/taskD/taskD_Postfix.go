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

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	elements := strings.Fields(input)

	stack := make([]int, 0)
	// var err error
	// var num int

	for _, element := range elements {
		if num, err := strconv.Atoi(element); err != nil {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			switch element {
			case "*":
				stack = append(stack, a*b)
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			}
		} else {
			stack = append(stack, num)
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, v := range stack {
		fmt.Fprintf(writer, "%d", v)
	}
	writer.Flush()
}
