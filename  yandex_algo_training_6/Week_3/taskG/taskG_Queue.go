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
	n, _ := strconv.Atoi(fstr[0])
	b, _ := strconv.Atoi(fstr[1])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}

	totalTime := 0
	queue := 0

	for i := 0; i < n; i++ {
		queue += array[i]
		totalTime += queue
		queue -= min(queue, b)
	}

	totalTime += queue

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, totalTime)
	writer.Flush()
}
