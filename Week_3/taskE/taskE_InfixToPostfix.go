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
	tokens, isValid := tokenize(input)
	if !isValid {
		fmt.Print("WRONG")
		return
	}

	array := make([]string, 0, len(tokens))
	stack := make([]string, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				array = append(array, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				fmt.Print("WRONG")
				return
			}
			stack = stack[:len(stack)-1]
		case "*", "/":
			for len(stack) > 0 && (stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/") {
				array = append(array, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "+", "-":
			for len(stack) > 0 && (stack[len(stack)-1] == "+" || stack[len(stack)-1] == "-" || stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/") {
				array = append(array, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		default:
			if _, err := strconv.Atoi(token); err != nil {
				fmt.Print("WRONG")
				return
			} else {
				array = append(array, token)
			}
		}
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			fmt.Print("WRONG")
			return
		}
		array = append(array, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	var err error
	var num int

	stack2 := make([]int, 0, len(array))

	for _, element := range array {
		if num, err = strconv.Atoi(element); err != nil {
			if len(stack2) < 2 {
				fmt.Print("WRONG")
				return
			}
			a := stack2[len(stack2)-2]
			b := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-2]

			switch element {
			case "*":
				stack2 = append(stack2, a*b)
			case "+":
				stack2 = append(stack2, a+b)
			case "-":
				stack2 = append(stack2, a-b)
			default:
				fmt.Print("WRONG")
				return
			}
		} else {
			stack2 = append(stack2, num)
		}
	}

	if len(stack2) != 1 {
		fmt.Print("WRONG")
		return
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(writer, "%d", stack2[0])
	writer.Flush()
}

func tokenize(input string) ([]string, bool) {
	var tokens []string
	var numBuilder strings.Builder
	var lastChar rune

	for _, char := range input {
		if char == ' ' {
			if numBuilder.Len() > 0 {
				tokens = append(tokens, numBuilder.String())
				numBuilder.Reset()
			}
			continue
		}
		if char >= '0' && char <= '9' {
			if lastChar == ' ' && numBuilder.Len() > 0 {
				return nil, false
			}
			numBuilder.WriteRune(char)
		} else {
			if numBuilder.Len() > 0 {
				tokens = append(tokens, numBuilder.String())
				numBuilder.Reset()
			}
			if char == '-' && (lastChar == ' ' || lastChar == '(' || lastChar == '+' || lastChar == '-' || lastChar == '*' || lastChar == '/') {
				numBuilder.WriteRune(char)
			} else {
				tokens = append(tokens, string(char))
			}
		}
		lastChar = char
	}
	if numBuilder.Len() > 0 {
		tokens = append(tokens, numBuilder.String())
	}
	return tokens, true
}
