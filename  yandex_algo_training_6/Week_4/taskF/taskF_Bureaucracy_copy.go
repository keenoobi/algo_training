package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fstr, _ := strconv.Atoi(strings.TrimSpace(input))

	input, _ = reader.ReadString('\n')
	scndstr := strings.Split(strings.TrimSpace(input), " ")

	// Инициализируем список смежности для представления дерева
	// Что такое список смежности?
	list := make([][]int, fstr+1)
	for i := 2; i <= fstr; i++ {
		parent, _ := strconv.Atoi(scndstr[i-2])
		list[parent] = append(list[parent], i)
	}
	fmt.Println(list)

	// Массивы для хранения глубины, размера поддерева и суммы глубин поддерева
	depth := make([]int, fstr+1)
	subtreeSize := make([]int, fstr+1)
	sumDepthSubtree := make([]int, fstr+1)

	// Рекурсивная функция для вычисления глубины и размера поддерева
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

	// Начинаем DFS с корневого сотрудника (Мирко, сотрудник 1)
	dfs(1)

	// Вычисляем общее количество монет для каждого сотрудника
	totalCoins := make([]int, fstr+1)
	for employee := 1; employee <= fstr; employee++ {
		totalCoins[employee] = sumDepthSubtree[employee] - subtreeSize[employee]*depth[employee] + subtreeSize[employee]
	}

	// Выводим общее количество монет для каждого сотрудника
	for employee := 1; employee <= fstr; employee++ {
		fmt.Printf("%d ", totalCoins[employee])
	}
	fmt.Println()
}
