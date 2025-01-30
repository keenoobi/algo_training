package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Ввод данных
	var n int
	fmt.Fscan(in, &n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &points[i].x, &points[i].y)
	}

	// Сортировка точек по x, затем по y
	// sort.Slice(points, func(i, j int) bool {
	// 	if points[i].x == points[j].x {
	// 		return points[i].y < points[j].y
	// 	}
	// 	return points[i].x < points[j].x
	// })

	// Жадный поиск троек
	used := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if used[j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if used[k] {
					continue
				}
				if isNonDegenerate(points[i], points[j], points[k]) {
					// Если тройка невырожденная, отмечаем её
					used[i], used[j], used[k] = true, true, true
					count++
					break
				}
			}
			if used[i] {
				break
			}
		}
	}

	// Вывод результата
	fmt.Fprintln(out, count)
}

// Проверка, что три точки не лежат на одной прямой
func isNonDegenerate(a, b, c Point) bool {
	det := a.x*(b.y-c.y) + b.x*(c.y-a.y) + c.x*(a.y-b.y)
	return det != 0
}
