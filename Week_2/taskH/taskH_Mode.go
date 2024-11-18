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
	totalEmployees := 0
	for i := 0; i < N; i++ {
		array[i], _ = strconv.Atoi(elements[i])
		totalEmployees += array[i]
	}

	halfTotal := (totalEmployees + 1) / 2
	currentEmployees := 0
	medianIndex := 0

	for i := 0; i < N; i++ {
		currentEmployees += array[i]
		if currentEmployees >= halfTotal {
			medianIndex = i
			break
		}
	}
	fmt.Println(medianIndex)

	minTransfers := 0
	for i := 0; i < N; i++ {
		minTransfers += abs(medianIndex-i) * array[i]
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, minTransfers)
	writer.Flush()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
