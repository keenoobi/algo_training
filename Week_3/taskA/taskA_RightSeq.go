package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type Stack []int

// func (s *Stack) push(v int) {
// 	*s = append(*s, v)
// }

// func (s *Stack) pop() (int, bool) {
// 	if len(*s) == 0 {
// 		return 0, false
// 	}
// 	index := len(*s) - 1
// 	element := (*s)[index]
// 	*s = (*s)[:index]
// 	return element, true
// }

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var stack []byte

	for i := 0; i < len(input); i++ {
		if input[i] == '(' || input[i] == '[' || input[i] == '{' {
			stack = append(stack, input[i])
		} else if len(stack) != 0 && ((input[i] == ')' && stack[len(stack)-1] == '(') || (input[i] == ']' && stack[len(stack)-1] == '[') || (input[i] == '}' && stack[len(stack)-1] == '{')) {
			stack = stack[:len(stack)-1]
		} else {
			fmt.Print("1No")
			return
		}
	}

	/* for i := 0; i < len(input); i++ {
		switch input[i] {
		case '(', '[', '{':
			stack = append(stack, input[i])
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				fmt.Print("no")
				return
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				fmt.Print("no")
				return
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				fmt.Print("no")
				return
			}
			stack = stack[:len(stack)-1]
		default:
			fmt.Print("no")
			return
		}
	} */

	if len(stack) == 0 {
		fmt.Print("yes")
	} else {
		fmt.Print("2no")
	}

}
