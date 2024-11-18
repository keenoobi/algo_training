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

	// Считываем количество улик
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)

	// Считываем веса улик
	weightsStr, _ := reader.ReadString('\n')
	weightsStr = strings.TrimSpace(weightsStr)
	weights := strings.Fields(weightsStr)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(weights[i])
	}

	// Считываем количество экспериментов и максимальное число перемещений
	mkStr, _ := reader.ReadString('\n')
	mkStr = strings.TrimSpace(mkStr)
	mk := strings.Fields(mkStr)
	m, _ := strconv.Atoi(mk[0])
	k, _ := strconv.Atoi(mk[1])

	// Считываем номера улик, с которых начинаются эксперименты
	startIndicesStr, _ := reader.ReadString('\n')
	startIndicesStr = strings.TrimSpace(startIndicesStr)
	startIndices := strings.Fields(startIndicesStr)

	// Массив для хранения результатов
	results := make([]int, n)

	// Предварительная обработка: заполняем массив результатов
	lastIndex := make(map[int]int) // Храним последний индекс для каждой весомости
	for i := 0; i < n; i++ {
		current := i
		count := 0

		// Перемещаемся влево по уликам
		for current > 0 && (array[current-1] < array[current] || (array[current-1] == array[current] && count < k)) {
			if array[current-1] == array[current] {
				count++
				if count > k {
					break
				}
			} else {
				count = 0
			}
			current--
		}
		results[i] = current + 1      // +1 для корректного индексирования (1-based)
		lastIndex[array[i]] = current // Обновляем последний индекс для текущей весомости
	}

	// Обрабатываем запросы и выводим результаты
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < m; i++ {
		xi, _ := strconv.Atoi(startIndices[i])
		fmt.Fprint(writer, results[xi-1], " ") // -1 для корректного индексирования (1-based)
	}
	writer.Flush()
}
