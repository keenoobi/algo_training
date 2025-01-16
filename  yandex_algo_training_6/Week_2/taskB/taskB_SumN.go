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

	prefixSums := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		prefixSums[i] = prefixSums[i-1] + array[i-1]
	}

	// count := make(map[int]int)
	// count[0] = 1
	// result := 0

	// for i := 1; i <= N; i++ {
	// 	if value, exist := count[prefixSums[i]-K]; exist {
	// 		result += value
	// 	}
	// 	count[prefixSums[i]]++
	// }

	// best solution
	var left, right, result, sum int

	for ; right < N; right++ {
		sum += array[right]
		for sum > K {
			sum -= array[left]
			left++
		}
		if sum == K {
			result++
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, result)
	writer.Flush()

	// // fmt.Print(array)

}
