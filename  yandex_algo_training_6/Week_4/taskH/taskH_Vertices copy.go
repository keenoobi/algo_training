package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(v, p int, cost []int, adj [][]int, dp1, dp0, parent []int) {
	dp1[v] = cost[v]
	dp0[v] = 0
	parent[v] = p

	for _, to := range adj[v] {
		if to == p {
			continue
		}
		dfs(to, v, cost, adj, dp1, dp0, parent)
		dp1[v] += min(dp1[to], dp0[to])
		dp0[v] += dp1[to]
	}
}

func findMarked(v int, parent []int, adj [][]int, dp1, dp0 []int, marked []bool) {
	if dp1[v] < dp0[v] {
		marked[v] = true
	}
	for _, to := range adj[v] {
		if to == parent[v] {
			continue
		}
		if marked[v] {
			findMarked(to, parent, adj, dp1, dp0, marked)
		} else {
			if dp1[to] < dp0[to] {
				findMarked(to, parent, adj, dp1, dp0, marked)
			} else {
				marked[to] = true
				findMarked(to, parent, adj, dp1, dp0, marked)
			}
		}
	}
}

func test() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	cost := make([]int, n)
	adj := make([][]int, n)
	dp1 := make([]int, n)
	dp0 := make([]int, n)
	parent := make([]int, n)

	for i := 0; i < n-1; i++ {
		edgeStr, _ := reader.ReadString('\n')
		edgeParts := strings.Split(strings.TrimSpace(edgeStr), " ")
		u, _ := strconv.Atoi(edgeParts[0])
		v, _ := strconv.Atoi(edgeParts[1])
		u--
		v--
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	costStr, _ := reader.ReadString('\n')
	costParts := strings.Split(strings.TrimSpace(costStr), " ")
	for i := 0; i < n; i++ {
		cost[i], _ = strconv.Atoi(costParts[i])
	}

	dfs(0, -1, cost, adj, dp1, dp0, parent)

	minCost := min(dp1[0], dp0[0])

	marked := make([]bool, n)
	findMarked(0, parent, adj, dp1, dp0, marked)

	markedVertices := make([]int, 0)
	for i := 0; i < n; i++ {
		if marked[i] {
			markedVertices = append(markedVertices, i+1)
		}
	}
	fmt.Println(minCost, len(markedVertices))
	fmt.Println(strings.Trim(fmt.Sprint(markedVertices), "[]"))
}
