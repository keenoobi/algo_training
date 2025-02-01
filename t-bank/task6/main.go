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

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &points[i].x, &points[i].y)
	}

	if allCollinear(points) {
		fmt.Println(0)
	} else {
		fmt.Println(n / 3)
	}
}

func allCollinear(points []Point) bool {
	if len(points) < 3 {
		return true
	}

	// Проверяем, коллинеарны ли первые три точки
	a, b, c := points[0], points[1], points[2]
	if !isCollinear(a, b, c) {
		return false
	}

	// Если первые три коллинеарны, проверяем остальные
	for i := 3; i < len(points); i++ {
		if !isCollinear(a, b, points[i]) {
			return false
		}
	}
	return true
}

func isCollinear(a, b, c Point) bool {
	// Вектор AB
	abx := b.x - a.x
	aby := b.y - a.y

	// Вектор AC
	acx := c.x - a.x
	acy := c.y - a.y

	// Определитель матрицы 2x2 (ABx * ACy - ABy * ACx)
	det := abx*acy - aby*acx
	return det == 0
}
*/
