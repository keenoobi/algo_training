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
	fstr := strings.Fields(input)
	N, _ := strconv.Atoi(fstr[0])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}

	stack := make([]int, 0, N)
	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[i] = -1
	}

	for i := 0; i < N; i++ {
		for len(stack) > 0 && array[i] < array[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, v := range result {
		fmt.Fprintf(writer, "%d ", v)
	}
	writer.Flush()
}
