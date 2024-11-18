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
	K, _ := strconv.Atoi(fstr[1])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}
	left := 0
	right := 1
	var result int

	// for left < N-1 {
	// 	for right < N && array[right]-array[left] <= K {
	// 		right++
	// 	}
	// 	if right < N {
	// 		result += N - right
	// 	}
	// 	left++
	// }

	for left < N && right < N {
		for right < N && array[right]-array[left] <= K {
			right++
		}
		result += N - right
		left++
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, result)
	writer.Flush()
}
