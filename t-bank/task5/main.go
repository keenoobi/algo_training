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

	// Чтение данных
	var n int
	var s int64
	fmt.Fscan(in, &n, &s)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Префиксные суммы
	pref := make([]int64, n)
	pref[0] = a[0]
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + a[i]
	}

	// Функция для вычисления суммы на подотрезке [l, r]
	getSum := func(l, r int) int64 {
		if l == 0 {
			return pref[r]
		}
		return pref[r] - pref[l-1]
	}

	// Динамическое программирование
	dp := make([]int64, n)
	result := int64(0)
	for i := n - 1; i >= 0; i-- {
		left := i
		right := n - 1
		if getSum(i, n-1) <= s {
			dp[i] = int64(n - i)
			result += dp[i]
			continue
		}

		// Бинарный поиск
		for left < right {
			mid := left + (right-left)/2
			if getSum(i, mid) <= s {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// left - первый индекс, где сумма на [i, left] > s
		dp[i] = dp[left] + int64(n-i)
		result += dp[i]
	}

	// Вывод результата
	fmt.Fprintln(out, result)
}

/*
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

		// Чтение данных
		var n, s int
		fmt.Fscan(in, &n, &s)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		// Вычисление ответа
		var result int = 0

		for left := 0; left < n; left++ {
			var currSum int = 0
			var parts int = 0
			for right := left; right < n; right++ {
				currSum += a[right]
				if currSum > s {
					parts++
					currSum = a[right]
				}
				result += parts + 1
			}
		}

		// Вывод результата
		fmt.Fprintln(out, result)
	}
*/

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

	// Чтение данных
	var n, s int
	fmt.Fscan(in, &n, &s)
	a := make([]int, n)

	// Ввод массива
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Создание массива префиксных сумм
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + a[i-1]
	}
	fmt.Fprintln(out, prefixSum)

	// Вычисление результата
	var result int = 0
	for left := 0; left < n; left++ {
		var parts int = 0
		currSum := 0
		for right := left; right < n; right++ {
			// Сумма отрезка [left, right] через префиксные суммы
			currSum = prefixSum[right+1] - prefixSum[left]
			if currSum > s {
				parts++
				// currSum = a[right] // начинаем новый кусок с текущего элемента
			}
			result += parts + 1
		}
	}

	// Вывод результата
	fmt.Fprintln(out, result)
} */
