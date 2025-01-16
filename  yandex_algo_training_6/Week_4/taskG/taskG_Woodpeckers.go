package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Чтение данных из stdin
func readInput() (int, int, int, [][2]int) {
	reader := bufio.NewReader(os.Stdin)
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return n, m, k, edges
}

// Функция проверки двухдольности графа с помощью DFS
func isBipartite(n int, edges [][2]int) (bool, []int) {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	colors := make([]int, n+1) // 0 - не посещён, 1 - первый цвет, 2 - второй цвет
	var dfs func(int, int) bool
	dfs = func(node, color int) bool {
		colors[node] = color
		for _, neighbor := range graph[node] {
			if colors[neighbor] == 0 {
				if !dfs(neighbor, 3-color) {
					return false
				}
			} else if colors[neighbor] == colors[node] {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if colors[i] == 0 {
			if !dfs(i, 1) {
				return false, nil
			}
		}
	}
	return true, colors
}

// Проверка на пересечение рёбер
func hasEdgeIntersections(edges [][2]int) bool {
	// Сортируем рёбра по первому элементу, а затем проверяем порядок второго
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0] == edges[j][0] {
			return edges[i][1] < edges[j][1]
		}
		return edges[i][0] < edges[j][0]
	})

	active := []int{}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if u > v {
			u, v = v, u // Переводим в упорядоченную пару
		}

		// Проверяем, есть ли пересечения
		for _, a := range active {
			if a > v {
				return true
			}
		}

		active = append(active, v)
	}

	return false
}

// Подсчёт количества размещений
func countPlacements(n int, k int, colors []int) int {
	modExp := func(base, exp, mod int) int {
		result := 1
		for exp > 0 {
			if exp%2 == 1 {
				result = (result * base) % mod
			}
			base = (base * base) % mod
			exp /= 2
		}
		return result
	}

	// Подсчёт числа вершин в каждом из двух цветов
	count1, count2 := 0, 0
	for i := 1; i <= n; i++ {
		if colors[i] == 1 {
			count1++
		} else if colors[i] == 2 {
			count2++
		}
	}

	// Количество размещений — это 2^(count1 + count2) по модулю k
	return modExp(2, count1+count2, k)
}

func main() {
	n, _, k, edges := readInput()

	// Проверяем, является ли граф двухдольным
	isBipartiteGraph, colors := isBipartite(n, edges)
	if !isBipartiteGraph {
		fmt.Println(0)
		return
	}

	// Проверяем, есть ли пересечения рёбер
	if hasEdgeIntersections(edges) {
		fmt.Println(0)
		return
	}

	// Если граф двухдольный, считаем количество размещений
	result := countPlacements(n, k, colors)
	fmt.Println(result)
}
