package main

import (
	"bufio"
	"container/list"
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
	k, _ := strconv.Atoi(fstr[1])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}

	deque := list.New()
	result := make([]int, n-k+1)

	for i := 0; i < n; i++ {
		if deque.Len() > 0 && deque.Front().Value.(int) <= i-k {
			deque.Remove(deque.Front())
		}

		for deque.Len() > 0 && array[i] <= array[deque.Back().Value.(int)] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)

		if i >= k-1 {
			result[i-k+1] = array[deque.Front().Value.(int)]
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, v := range result {
		fmt.Fprintf(writer, "%d\n", v)
	}
	writer.Flush()
}
