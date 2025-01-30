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

// /* package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Буферизованный ввод и вывод
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var n, x, y, z int
// 	fmt.Fscan(in, &n, &x, &y, &z) // Считываем n, x, y, z

// 	array := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &array[i]) // Считываем последовательность
// 	}

// 	// Функция для вычисления операций
// 	ops := func(num int, div int) int {
// 		mod := num % div
// 		if mod == 0 {
// 			return 0
// 		}
// 		return div - mod
// 	}

// 	// Находим минимальные операции для каждого делителя
// 	minX := int(1 << 30)
// 	minY := int(1 << 30)
// 	minZ := int(1 << 30)

// 	var opsX, opsY, opsZ int

// 	for _, num := range array {
// 		opsX = ops(num, x)
// 		if opsX < minX {
// 			minX = opsX
// 		}
// 		opsY = ops(num, y)
// 		if opsY < minY {
// 			minY = opsY
// 		}
// 		opsZ = ops(num, z)
// 		if opsZ < minZ {
// 			minZ = opsZ
// 		}
// 	}

// 	var minOps int
// 	if minX == minY && minX == minZ {
// 		minOps = minX
// 	} else if minX == minY {
// 		minOps = minX + minZ
// 	} else if minX == minZ {
// 		minOps = minX + minY
// 	} else if minY == minZ {
// 		minOps = minX + minY
// 	} else {
// 		minOps = minX + minY + minZ
// 	}

// 	// Выводим сумму минимальных операций
// 	fmt.Fprintln(out, minOps)
// }
// */

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var n, x, y, z int
// 	fmt.Fscan(in, &n, &x, &y, &z)

// 	array := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &array[i])
// 	}

// 	// Функция для вычисления операций до кратности div
// 	ops := func(num, div int) int {
// 		mod := num % div
// 		if mod == 0 {
// 			return 0
// 		}
// 		return div - mod
// 	}

// 	// Минимальные операции для каждого делителя
// 	minX, minY, minZ := int(1e9), int(1e9), int(1e9)
// 	// Минимальные операции для комбинаций кратностей
// 	minXY, minXZ, minYZ := int(1e9), int(1e9), int(1e9)
// 	minXYZ := int(1e9)

// 	idX := 0
// 	idY := 0
// 	idZ := 0

// 	for i, num := range array {
// 		opX := ops(num, x)
// 		opY := ops(num, y)
// 		opZ := ops(num, z)
// 		opXY := ops(num, lcm(x, y))
// 		opXZ := ops(num, lcm(x, z))
// 		opYZ := ops(num, lcm(y, z))
// 		opXYZ := ops(num, lcm3(x, y, z))

// 		// Отдельные минимумы
// 		if opX < minX {
// 			minX = opX
// 			idX = i
// 		}
// 		if opY < minY {
// 			minY = opY
// 			idY = i
// 		}
// 		if opZ < minZ {
// 			minZ = opZ
// 			idZ = i
// 		}

// 		// Комбинации
// 		if opXY < minXY {
// 			minXY = opXY
// 		}
// 		if opXZ < minXZ {
// 			minXZ = opXZ
// 		}
// 		if opYZ < minYZ {
// 			minYZ = opYZ
// 		}

// 		// Все три сразу
// 		if opXYZ < minXYZ {
// 			minXYZ = opXYZ
// 		}

// 	}

// 	// fmt.Fprintln(out, minXY)
// 	// fmt.Fprintln(out, minXZ)
// 	// fmt.Fprintln(out, minYZ)
// 	// fmt.Fprintln(out, minXYZ)

// 	if minX+minY < minXY && idX != idY {
// 		minXY = minX + minY
// 	}
// 	if minX+minZ < minXZ && idX != idZ {
// 		minXZ = minX + minZ
// 	}
// 	if minY+minZ < minYZ && idY != idZ {
// 		minYZ = minY + minZ
// 	}

// 	if minX+minYZ < minXYZ && idX != idY {
// 		minXYZ = minX + minYZ
// 	}
// 	if minZ+minXY < minXYZ && idX != idZ {
// 		minXYZ = minZ + minXY
// 	}
// 	if minY+minXZ < minXYZ && idY != idZ {
// 		minXYZ = minY + minXZ
// 	}

// 	fmt.Fprintln(out, minXYZ)
// }

// // Функция для нахождения наибольшего общего делителя (НОД)
// func gcd(a, b int) int {
// 	for b != 0 {
// 		a, b = b, a%b // Обновляем a и b
// 	}
// 	return a // Возвращаем НОД
// }

// // Функция для нахождения наименьшего общего кратного (НОК)
// func lcm(a, b int) int {
// 	return (a * b) / gcd(a, b) // Используем формулу для НОК
// }

// // Функция для нахождения наименьшего общего кратного (НОК) для трех чисел
// func lcm3(a, b, c int) int {
// 	return (a * b / gcd(a, b)) * c / gcd((a*b/gcd(a, b)), c)
// }

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

// Функция нахождения НОК через НОД
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	var n, x, y, z int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &n, &x, &y, &z)
	defer out.Flush()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	Lxy := lcm(x, y)
	Lxz := lcm(x, z)
	Lyz := lcm(y, z)
	Lxyz := lcm(Lxy, z)

	// Инициализируем переменные большим числом
	const inf = 1 << 30
	cx, cy, cz := inf, inf, inf
	cxy, cxz, cyz, cxyz := inf, inf, inf, inf

	for _, a := range arr {
		dx := (x - a%x) % x
		dy := (y - a%y) % y
		dz := (z - a%z) % z
		dxy := (Lxy - a%Lxy) % Lxy
		dxz := (Lxz - a%Lxz) % Lxz
		dyz := (Lyz - a%Lyz) % Lyz
		dxyz := (Lxyz - a%Lxyz) % Lxyz

		// Обновляем минимальные значения
		cx = min(cx, dx)
		cy = min(cy, dy)
		cz = min(cz, dz)
		cxy = min(cxy, dxy)
		cxz = min(cxz, dxz)
		cyz = min(cyz, dyz)
		cxyz = min(cxyz, dxyz)
	}

	// Берем минимум из всех возможных комбинаций
	result := min(cx+cy+cz, cxy+cz, cxz+cy, cyz+cx, cxyz)
	fmt.Fprintln(out, cx+cy+cz, cxy+cz, cxz+cy, cyz+cx, cxyz)

	fmt.Fprint(out, result)
}

func min(a ...int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
} */
