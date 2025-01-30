package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	op := make([]int, 0, n-2)
	for i := 2; i < n; i++ {
		if a[i] < a[0] {
			op = append(op, a[0]-a[i])
		} else if a[i] > a[1] {
			op = append(op, a[i]-a[1])
		} else {
			op = append(op, 0)
		}
	}

	sort.Ints(op)

	result := 0
	for i := 0; i < m; i++ {
		result += op[i]
	}

	fmt.Fprintln(out, result)
}

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Читаем входные данные
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	// Корректируем a1 и a2, если необходимо
	l, r := a[0], a[1]
	if l > r {
		l, r = r, l
	}

	// Собираем оставшиеся элементы в массив b
	b := make([]int, n-2)
	copy(b, a[2:])

	// Сортируем массив b
	sort.Ints(b)

	// Минимальные корректировки
	minChanges := int(1e9)

	// Проходим по массиву b "окном" размера m
	for i := 0; i <= len(b)-m; i++ {
		left, right := b[i], b[i+m-1]

		// Вычисляем корректировки для включения [left, right] в [l, r]
		adjustL := max(0, l-left)
		adjustR := max(0, right-r)
		minChanges = min(minChanges, adjustL+adjustR)
	}

	// Выводим результат
	fmt.Println(minChanges)
} */
