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
	n, _ := strconv.Atoi(input)

	stack := make([]int, 0)
	prefixSums := make([]int, 0)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if line[0] == '+' || line[0] == '-' || line[0] == '?' {
			command := line[0]
			number := 0

			if len(line) > 1 {
				number, _ = strconv.Atoi(line[1:])
			}
			switch command {
			case '+':
				stack = append(stack, number)
				if len(prefixSums) == 0 {
					prefixSums = append(prefixSums, number)
				} else {
					prefixSums = append(prefixSums, prefixSums[len(prefixSums)-1]+number)
				}
			case '-':
				fmt.Println(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
				prefixSums = prefixSums[:len(prefixSums)-1]
			case '?':
				if number >= len(prefixSums) {
					fmt.Println(prefixSums[len(prefixSums)-1])
				} else {
					fmt.Println(prefixSums[len(prefixSums)-1] - prefixSums[len(prefixSums)-1-number])
				}
			}
		}
	}
}
