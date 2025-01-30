/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Чтение количества дней
	var n int
	fmt.Fscan(in, &n)

	// Чтение бюджетов
	budgets := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &budgets[i])
	}

	// Предварительное вычисление степеней двойки
	flowers := make([]int64, 61) // 2^0 до 2^60
	for i := 0; i <= 60; i++ {
		flowers[i] = 1 << i
	}

	// Обработка каждого дня
	for _, budget := range budgets {
		// Если бюджет меньше 7, букет невозможен
		if budget < 7 {
			fmt.Fprintln(out, -1)
			continue
		}

		// Находим максимальную сумму трёх степеней двойки
		maxSum := findMaxSum(budget, flowers)
		fmt.Fprintln(out, maxSum)
	}
}

// Функция для нахождения максимальной суммы трёх степеней двойки
func findMaxSum(budget int64, flowers []int64) int64 {
	maxSum := int64(-1)

	// Используем два указателя для поиска максимальной суммы
	for i := 0; i < len(flowers)-2; i++ {
		left, right := i+1, len(flowers)-1
		for left < right {
			sum := flowers[i] + flowers[left] + flowers[right]
			if sum <= budget {
				if sum > maxSum {
					maxSum = sum
				}
				left++
			} else {
				right--
			}
		}
	}

	return maxSum
} */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	budgets := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &budgets[i])
	}

	flowers := make([]int64, 61) // 2^0 до 2^60
	for i := 0; i <= 60; i++ {
		flowers[i] = 1 << i
	}

	for _, budget := range budgets {
		if budget < 7 {
			fmt.Fprintln(out, -1)
			continue
		}

		maxSum := int64(-1)

		for i := 0; i < len(flowers)-2; i++ {
			left, right := i+1, len(flowers)-1
			for left < right {
				sum := flowers[i] + flowers[left] + flowers[right]
				if sum <= budget {
					if sum > maxSum {
						maxSum = sum
						fmt.Fprintln(out, flowers[i], flowers[left], flowers[right])
					}
					left++
				} else {
					right--
				}
			}
		}

		fmt.Fprintln(out, maxSum)
	}
}

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(reader, &n)

	budgets := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &budgets[i])
	}

	results := make([]int64, n)

	for day := 0; day < n; day++ {
		budget := budgets[day]
		if budget < 7 {
			results[day] = -1
			continue
		}

		maxCost := int64(0)
		for i := 0; i < 100; i++ {
			for j := i + 1; j < 100; j++ {
				for k := j + 1; k < 100; k++ {
					cost := (1 << i) + (1 << j) + (1 << k) // Эквивалентно 2^i + 2^j + 2^k
					if int64(cost) <= budget {
						if int64(cost) > maxCost {
							maxCost = int64(cost)
						}
					}
				}
			}
		}

		results[day] = maxCost
	}

	for _, cost := range results {
		fmt.Fprintln(writer, cost)
	}

	writer.Flush() // Сбрасываем буфер и выводим данные
} */
