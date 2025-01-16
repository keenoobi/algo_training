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

	fstr, _ := reader.ReadString('\n')
	V, _ := strconv.Atoi(strings.TrimSpace(fstr))

	list := make([][]int, V+1)

	for i := 0; i < V-1; i++ {
		nextline, _ := reader.ReadString('\n')
		edges := strings.Split(strings.TrimSpace(nextline), " ")

		edge1, _ := strconv.Atoi(edges[0])
		edge2, _ := strconv.Atoi(edges[1])

		list[edge1] = append(list[edge1], edge2)
		list[edge2] = append(list[edge2], edge1)
	}

	size := make([]int, V+1)

	var DFS func(int, int)
	DFS = func(currentV, parentV int) {
		size[currentV] = 1
		for _, neighbor := range list[currentV] {
			if neighbor != parentV {
				DFS(neighbor, currentV)
				size[currentV] += size[neighbor]
			}
		}
	}

	DFS(1, -1)

	for i := 1; i <= V; i++ {
		fmt.Printf("%d ", size[i])
	}
}
