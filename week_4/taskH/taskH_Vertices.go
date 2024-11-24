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

	N, _ := strconv.Atoi(strings.TrimSpace(input))

	edges := make([][]int, N)

	for i := 0; i < N-1; i++ {
		input, _ = reader.ReadString('\n')
		str := strings.Split(strings.TrimSpace(input), " ")

		u, _ := strconv.Atoi(str[0])
		u--
		v, _ := strconv.Atoi(str[1])
		v--

		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)

	}
	fmt.Println(edges)

	input, _ = reader.ReadString('\n')
	str := strings.Split(strings.TrimSpace(input), " ")
	costs := make([]int, N)

	for i := 0; i < N; i++ {
		costs[i], _ = strconv.Atoi(str[i])
	}

	dp0 := make([]int, N)
	dp1 := make([]int, N)

	fmt.Println(costs)
}
