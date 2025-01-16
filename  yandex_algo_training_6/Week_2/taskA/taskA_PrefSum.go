package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	elements := strings.Fields(scanner.Text())

	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i], _ = strconv.Atoi(elements[i])
	}

	prefixSums := make([]int, n)
	prefixSums[0] = sequence[0]
	for i := 1; i < n; i++ {
		prefixSums[i] = prefixSums[i-1] + sequence[i]
	}

	for _, sum := range prefixSums {
		fmt.Print(sum, " ")
	}

	/* 	prefixSums := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		prefixSums[i] = prefixSums[i-1] + sequence[i-1]
	}

	for _, sum := range prefixSums[1:] {
		fmt.Print(sum, " ")
	} */

}
