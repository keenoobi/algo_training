package main

import (
	"bufio"
	"fmt"
	"os"
)

// НОК двух чисел
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}

// НОД двух чисел
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// НОК трёх чисел
func lcmThree(a, b, c int64) int64 {
	return lcm(a, lcm(b, c))
}

// Минимальное количество операций
func getOps(a, target int64) int64 {
	return (target - a%target) % target
}

func main() {
	// Чтение входных данных
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var x, y, z int64
	fmt.Fscan(in, &n, &x, &y, &z)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	if n == 1 {
		opsXYZ := getOps(a[0], lcmThree(x, y, z))
		fmt.Fprintln(out, opsXYZ)
		return
	}

	const inf int64 = 1 << 60 // Используем 2^60 как "бесконечность"
	cx, cy, cz := inf, inf, inf
	cxy, cxz, cyz, cxyz := inf, inf, inf, inf

	for _, ai := range a {
		// Вычисляем операции для текущего элемента
		opsX := getOps(ai, x)
		opsY := getOps(ai, y)
		opsZ := getOps(ai, z)
		opsXY := getOps(ai, lcm(x, y))
		opsXZ := getOps(ai, lcm(x, z))
		opsYZ := getOps(ai, lcm(y, z))
		opsXYZ := getOps(ai, lcmThree(x, y, z))

		// Обновляем переменные
		cx = min(cx, opsX)
		cy = min(cy, opsY)
		cz = min(cz, opsZ)
		cxy = min(cxy, cx+opsY, cy+opsX, opsXY)
		cxz = min(cxz, cx+opsZ, cz+opsX, opsXZ)
		cyz = min(cyz, cy+opsZ, cz+opsY, opsYZ)
		cxyz = min(cxyz, cxy+opsZ, cxz+opsY, cyz+opsX, opsXYZ)
	}

	// Ответ — минимальное количество операций
	result := min(cx+cy+cz, cxy+cz, cxz+cy, cyz+cx, cxyz)
	fmt.Fprintln(out, result)
}

// Функция для нахождения минимума среди нескольких чисел
func min(a ...int64) int64 {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return (a / gcd(a, b)) * b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var x, y, z int64
	fmt.Fscan(in, &n, &x, &y, &z)

	// Вычисляем LCM для всех пар и тройки
	lcmXY := lcm(x, y)
	lcmXZ := lcm(x, z)
	lcmYZ := lcm(y, z)
	lcmXYZ := lcm(lcmXY, z)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Инициализация минимальных значений
	minX := int64(1e18)
	minY := int64(1e18)
	minZ := int64(1e18)
	minXY := int64(1e18)
	minXZ := int64(1e18)
	minYZ := int64(1e18)
	minXYZ := int64(1e18)

	for _, num := range a {
		// Вычисление стоимости для x, y, z
		costX := (x - (num % x)) % x
		costY := (y - (num % y)) % y
		costZ := (z - (num % z)) % z

		// Обновление минимальных значений
		if costX < minX {
			minX = costX
		}
		if costY < minY {
			minY = costY
		}
		if costZ < minZ {
			minZ = costZ
		}

		// Вычисление стоимости для пар и тройки
		costXY := (lcmXY - (num % lcmXY)) % lcmXY
		if costXY < minXY {
			minXY = costXY
		}

		costXZ := (lcmXZ - (num % lcmXZ)) % lcmXZ
		if costXZ < minXZ {
			minXZ = costXZ
		}

		costYZ := (lcmYZ - (num % lcmYZ)) % lcmYZ
		if costYZ < minYZ {
			minYZ = costYZ
		}

		costXYZ := (lcmXYZ - (num % lcmXYZ)) % lcmXYZ
		if costXYZ < minXYZ {
			minXYZ = costXYZ
		}
	}

	// Вычисление всех возможных комбинаций
	option1 := minX + minY + minZ
	option2 := minXY + minZ
	option3 := minXZ + minY
	option4 := minYZ + minX
	option5 := minXYZ

	// Находим минимальный вариант
	minOption := option1
	if option2 < minOption {
		minOption = option2
	}
	if option3 < minOption {
		minOption = option3
	}
	if option4 < minOption {
		minOption = option4
	}
	if option5 < minOption {
		minOption = option5
	}

	// Учитываем случай, когда n < 3
	if n < 2 {
		minOption = minXYZ
	}

	fmt.Fprintln(out, minOption)
} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return (a / gcd(a, b)) * b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var x, y, z int64
	fmt.Fscan(in, &n, &x, &y, &z)

	// Вычисляем LCM для всех пар и тройки
	lcmXY := lcm(x, y)
	lcmXZ := lcm(x, z)
	lcmYZ := lcm(y, z)
	lcmXYZ := lcm(lcmXY, z)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Инициализация минимальных значений
	minX := int64(1e18)
	minY := int64(1e18)
	minZ := int64(1e18)
	minXY := int64(1e18)
	minXZ := int64(1e18)
	minYZ := int64(1e18)
	minXYZ := int64(1e18)

	for _, num := range a {
		// Вычисление стоимости для x, y, z
		costX := (x - (num % x)) % x
		costY := (y - (num % y)) % y
		costZ := (z - (num % z)) % z

		// Обновление минимальных значений
		if costX < minX {
			minX = costX
		}
		if costY < minY {
			minY = costY
		}
		if costZ < minZ {
			minZ = costZ
		}

		// Вычисление стоимости для пар и тройки
		costXY := (lcmXY - (num % lcmXY)) % lcmXY
		if costXY < minXY {
			minXY = costXY
		}

		costXZ := (lcmXZ - (num % lcmXZ)) % lcmXZ
		if costXZ < minXZ {
			minXZ = costXZ
		}

		costYZ := (lcmYZ - (num % lcmYZ)) % lcmYZ
		if costYZ < minYZ {
			minYZ = costYZ
		}

		costXYZ := (lcmXYZ - (num % lcmXYZ)) % lcmXYZ
		if costXYZ < minXYZ {
			minXYZ = costXYZ
		}
	}

	// Вычисление всех возможных комбинаций
	option1 := minX + minY + minZ
	option2 := minXY + minZ
	option3 := minXZ + minY
	option4 := minYZ + minX
	option5 := minXYZ

	// Находим минимальный вариант
	minOption := option1
	if option2 < minOption {
		minOption = option2
	}
	if option3 < minOption {
		minOption = option3
	}
	if option4 < minOption {
		minOption = option4
	}
	if option5 < minOption {
		minOption = option5
	}

	fmt.Fprintln(out, minOption)
}
*/
