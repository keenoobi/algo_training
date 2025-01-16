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
	K, _ := strconv.Atoi(fstr[1])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}

	sort.Ints(array)
	fmt.Println(array)
	left, right, result := 0, 1, 1
	// done := false

	// for right < N {
	// 	for right < N && array[right]-array[left] <= K {
	// 		right++
	// 		done = true
	// 	}
	// 	if right < N && done {
	// 		result += 1
	// 		done = false
	// 	}

	// 	left = right
	// 	right++
	// }

	// wtf is that solution?

	/* 	found := 0

	   	for i, j := right, left; i < N; i++ {
	   		if array[i]-array[j] <= K {
	   			continue
	   		} else {
	   			j++
	   			found++
	   		}
	   	}
	   	result = N - found */

	for left < N && right < N {
		if array[right]-array[left] <= K {
			result = max(result, right-left+1)
			right++
		} else {
			left++
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, result)
	writer.Flush()
}
