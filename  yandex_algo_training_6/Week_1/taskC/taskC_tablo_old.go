package main

import (
	"fmt"
	"math"
	"strings"
)

func checkBounds(array [][]string, n int) (x1, y1, x2, y2 int) {
	x1, y1, x2, y2 = n, n, 0, 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if array[i][j] == "#" {
				x1 = int(math.Min(float64(x1), float64(i)))
				y1 = int(math.Min(float64(y1), float64(j)))
				x2 = int(math.Max(float64(x2), float64(i)))
				y2 = int(math.Max(float64(y2), float64(j)))
			}
		}
	}
	return
}

func checkRectangle(array [][]string, x1, y1, x2, y2 int) bool {
	for i := x1; i <= x2; i++ {
		for j := y1; j < y2; j++ {
			if array[i][j] != "#" {
				return false
			}
		}
	}
	return true
}

func checkI(array [][]string, x1, y1, x2, y2 int) bool {
	I := true
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if array[i][j] != "#" {
				I = false
			}
		}
	}
	if I && (x2-x1 >= y2-y1) {
		return true
	}
	return false
}

func checkO(array [][]string, x1, y1, x2, y2 int) bool {
	x3, y3, x4, y4 := -1, -1, -1, -1

	for i := x1 + 1; i < x2; i++ {
		for j := y1 + 1; j < y2; j++ {
			if array[i][j] == "." {
				if x3 == -1 || i < x3 {
					x3 = i
				}
				if y3 == -1 || j < y3 {
					y3 = j
				}
				if i > x4 {
					x4 = i
				}
				if j > y4 {
					y4 = j
				}
			}
		}
	}

	if x3 > x1 && y3 > y1 && x4 < x2 && y4 < y2 {
		// Дополнительно проверяем, что внешний прямоугольник состоит только из "#"
		for i := x1; i <= x2; i++ {
			if array[i][y1] != "#" || array[i][y2] != "#" {
				return false
			}
		}
		for j := y1; j <= y2; j++ {
			if array[x1][j] != "#" || array[x2][j] != "#" {
				return false
			}
		}
		return true
	}

	return false
}

func checkC(array [][]string, x1, y1, x2, y2 int) bool {
	// Проверяем, что верхняя, нижняя и левая границы прямоугольника заполнены "#"
	for i := x1; i <= x2; i++ {
		if array[i][y1] != "#" {
			return false
		}
	}
	for j := y1; j <= y2; j++ {
		if array[x1][j] != "#" || array[x2][j] != "#" {
			return false
		}
	}
	// Поиск внутреннего прямоугольника, который касается правой границы внешнего прямоугольника
	var x3, y3, x4, y4 int
	found := false

	for i := x1 + 1; i < x2 && !found; i++ {
		for j := y1 + 1; j <= y2 && !found; j++ {
			if array[i][j] == "." {
				// Нашли верхний левый угол внутреннего прямоугольника
				x3, y3 = i, j
				found = true
			}
		}

	}

	if !found {
		return false
	}

	x4 = x3
	for i := x3 + 1; i < x2 && array[i][y3] == "."; i++ {
		x4 = i
	}

	y4 = y3
	for j := y3 + 1; j <= y2 && array[x3][j] == "."; j++ {
		y4 = j
	}

	// fmt.Printf("x1 = %d\n", x1)
	// fmt.Printf("y1 = %d\n", y1)
	// fmt.Printf("x2 = %d\n", x2)
	// fmt.Printf("y2 = %d\n", y2)
	// fmt.Printf("x3 = %d\n", x3)
	// fmt.Printf("y3 = %d\n", y3)
	// fmt.Printf("x4 = %d\n", x4)
	// fmt.Printf("y4 = %d\n", y4)

	if x3 > x1 && y3 > y1 && (x3 < x4 || x3 == x4) && (y3 < y4 || y3 == y4) && (y3 < y2 || y3 == y2) && y4 == y2 {
		// Дополнительно проверяем, что внешний прямоугольник состоит только из "#"
		for i := x3; i <= x4; i++ {
			if array[i][y3] != "." || array[i][y4] != "." {
				return false
			}
		}
		for j := y3; j <= y4; j++ {
			if array[x3][j] != "." || array[x4][j] != "." {
				return false
			}
		}
		return true
	}

	return false
}

func checkL(array [][]string, x1, y1, x2, y2 int) bool {
	fmt.Print(x1, y1, x2, y2)

	for i := x1; i <= x2; i++ {
		if array[i][y1] != "#" {
			return false
		}
	}
	for j := y1; j <= y2; j++ {
		if array[x2][j] != "#" {
			return false
		}
	}

	var x3, y3, x4, y4 int
	found := false

	for i := x1; i < x2 && !found; i++ {
		for j := y1 + 1; j <= y2 && !found; j++ {
			if array[i][j] == "." {
				x3, y3 = i, j
				found = true
			}
		}

	}

	if !found {
		return false
	}

	x4 = x3
	for i := x3 + 1; i < x2 && array[i][y3] == "."; i++ {
		x4 = i
	}

	y4 = y3
	for j := y3 + 1; j <= y2 && array[x3][j] == "."; j++ {
		y4 = j
	}

	if x3 == x1 && y3 > y1 && y4 == y2 && x4 < x2 {
		for i := x3; i <= x4; i++ {
			if array[i][y3] != "." || array[i][y4] != "." {
				return false
			}
		}
		for j := y3; j <= y4; j++ {
			if array[x3][j] != "." || array[x4][j] != "." {
				return false
			}
		}
		return true
	}

	return false
}

func old() {
	var n int
	var line string

	fmt.Scan(&n)

	array := make([][]string, n)
	// array := [][]string{}
	for i := 0; i < n; i++ {
		fmt.Scan(&line)

		elements := strings.Split(line, "")

		array[i] = elements
	}

	x1, y1, x2, y2 := checkBounds(array, n)

	switch {
	case checkI(array, x1, y1, x2, y2):
		fmt.Println("I")
	case checkO(array, x1, y1, x2, y2):
		fmt.Println("O")
	case checkC(array, x1, y1, x2, y2):
		fmt.Println("C")
	case checkL(array, x1, y1, x2, y2):
		fmt.Println("L")
	// case checkH(array, x1, y1, x2, y2):
	// 	fmt.Println("H")
	// case checkP(array, x1, y1, x2, y2):
	// 	fmt.Println("P")
	default:
		fmt.Println("X")
	}

}
