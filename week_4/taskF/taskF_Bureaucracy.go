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
	fstr, _ := strconv.Atoi(strings.TrimSpace(input))

	input, _ = reader.ReadString('\n')
	scndstr := strings.Split(strings.TrimSpace(input), " ")

	list := make([][]int, fstr+1)
	for i := 2; i <= fstr; i++ {
		parent, _ := strconv.Atoi(scndstr[i-2])
		list[parent] = append(list[parent], i)
	}
	fmt.Println(list)

	depth := make([]int, fstr+1)
	subtreeSize := make([]int, fstr+1)
	sumDepthSubtree := make([]int, fstr+1)

	var dfs func(currentNode int)
	dfs = func(currentNode int) {
		subtreeSize[currentNode] = 1
		sumDepthSubtree[currentNode] = depth[currentNode]

		for _, subordinate := range list[currentNode] {
			depth[subordinate] = depth[currentNode] + 1
			dfs(subordinate)
			subtreeSize[currentNode] += subtreeSize[subordinate]
			sumDepthSubtree[currentNode] += sumDepthSubtree[subordinate]
		}
	}

	dfs(1)

	totalCoins := make([]int, fstr+1)
	for employee := 1; employee <= fstr; employee++ {
		totalCoins[employee] = sumDepthSubtree[employee] - subtreeSize[employee]*depth[employee] + subtreeSize[employee]
		fmt.Printf("%d ", totalCoins[employee])
	}

}
