package main

import (
	"fmt"
)

var numVertices int
var adjacencyList [][]int

func main() {
	fmt.Scan(&numVertices)
	adjacencyList = make([][]int, numVertices+1)
	for i := 0; i < numVertices-1; i++ {
		var vertex1, vertex2 int
		fmt.Scan(&vertex1, &vertex2)
		adjacencyList[vertex1] = append(adjacencyList[vertex1], vertex2)
		adjacencyList[vertex2] = append(adjacencyList[vertex2], vertex1)
	}

	// Функция для выполнения BFS и возврата самой дальней вершины и её расстояния от начальной вершины
	var BFS func(start int) (farthest int, distance int)
	BFS = func(start int) (farthest int, distance int) {
		visited := make([]bool, numVertices+1)
		queue := []int{start}
		visited[start] = true
		level := 0
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				current := queue[i]
				farthest = current
				for _, neighbor := range adjacencyList[current] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
			queue = queue[size:]
			level++
			distance = level - 1 // Вычитаем 1, так как уровень начинается с 1
		}
		return farthest, distance
	}

	// Находим диаметр дерева
	farthestNode1, _ := BFS(1)
	farthestNode2, treeDiameter := BFS(farthestNode1)

	// Находим путь для диаметра
	parent := make([]int, numVertices+1)
	var findPath func(current, target int)
	findPath = func(current, target int) {
		if current == target {
			return
		}
		for _, neighbor := range adjacencyList[current] {
			if neighbor != parent[current] {
				parent[neighbor] = current
				findPath(neighbor, target)
				if neighbor == target {
					return
				}
			}
		}
	}
	findPath(farthestNode1, farthestNode2)

	// Отмечаем вершины в пути диаметра
	visitedDiameter := make([]bool, numVertices+1)
	currentNode := farthestNode2
	visitedDiameter[currentNode] = true
	for currentNode != farthestNode1 {
		currentNode = parent[currentNode]
		visitedDiameter[currentNode] = true
	}

	// Находим самый длинный путь в оставшемся дереве
	var longestPath int
	for vertex := 1; vertex <= numVertices; vertex++ {
		if !visitedDiameter[vertex] {
			_, distance := BFS(vertex)
			if distance > longestPath {
				longestPath = distance
			}
		}
	}

	// Вычисляем произведение
	product := treeDiameter * longestPath
	fmt.Println(product)
}
