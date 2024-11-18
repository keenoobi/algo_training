package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fstr := strings.Fields(input)
	N, _ := strconv.Atoi(fstr[0])
	// K, _ := strconv.Atoi(fstr[1])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}

	sort.Ints(array)

	left, right := N/2-1, N/2
	result := make([]int, 0, N)

	for len(result) < N {
		if (N-len(result))%2 == 1 {
			result = append(result, array[right])
			right++
		} else {
			result = append(result, array[left])
			left--
		}
	}

	// array2 := make([]int, 0, N)

	// var medianIndex int

	// for len(array) > 0 {
	// 	if len(array)%2 != 0 {
	// 		medianIndex = len(array) / 2
	// 	} else {
	// 		medianIndex = len(array)/2 - 1
	// 	}
	// 	array2 = append(array2, array[medianIndex])
	// 	array = append(array[:medianIndex], array[medianIndex+1:]...)
	// }

	writer := bufio.NewWriter(os.Stdout)
	for _, v := range result {
		fmt.Fprintf(writer, "%d ", v)
	}
	writer.Flush()
}
