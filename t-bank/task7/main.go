package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = int(int64(res) * int64(a) % int64(mod))
		}
		a = int(int64(a) * int64(a) % int64(mod))
		b /= 2
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] %= MOD
	}

	// Вычисление степенных сумм
	S := make([]int, k+1)
	S[0] = n % MOD

	for _, x := range a {
		current := x
		for m := 1; m <= k; m++ {
			S[m] = (S[m] + current) % MOD
			current = int(int64(current) * int64(x) % MOD)
		}
	}

	// Предвычисление обратных элементов
	inv := make([]int, k+1)
	inv[1] = 1
	for m := 2; m <= k; m++ {
		inv[m] = int(int64(MOD-MOD/m) * int64(inv[MOD%m]) % MOD)
	}

	// Вычисление степеней 2
	pow2 := make([]int, k+1)
	pow2[0] = 1
	for p := 1; p <= k; p++ {
		pow2[p] = int(int64(pow2[p-1]) * 2 % MOD)
	}

	inv2 := powMod(2, MOD-2, MOD)

	// Вычисление f(p) для каждого p от 1 до k
	comb := make([]int, k+1)
	for p := 1; p <= k; p++ {
		comb[0] = 1
		for m := 1; m <= p; m++ {
			comb[m] = int(int64(comb[m-1]) * int64(p-m+1) % MOD * int64(inv[m]) % MOD)
		}

		sumPart := 0
		for m := 0; m <= p; m++ {
			sumPart = (sumPart + comb[m]*S[m]%MOD*S[p-m]%MOD) % MOD
		}

		term := ((sumPart-pow2[p]*S[p]%MOD)%MOD + MOD) % MOD
		res := term * inv2 % MOD

		fmt.Fprintln(out, res)
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// )

// const mod = 998244353

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	// Чтение входных данных
// 	var n, k int
// 	fmt.Fscan(in, &n, &k)

// 	array := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &array[i])
// 	}
// 	var sum []int
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			sum = append(sum, array[i]+array[j])
// 		}
// 	}
// 	result := 0
// 	for p := 1; p <= k; p++ {
// 		for _, elem := range sum {
// 			result += int(math.Pow(float64(elem), float64(p))) % mod
// 		}
// 		fmt.Println(result)
// 		result = 0
// 	}

// }

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

// Быстрое возведение в степень по модулю
func fastPow(base, exp, mod int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp /= 2
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Чтение входных данных
	var n, k int
	fmt.Fscan(in, &n, &k)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &array[i])
	}

	// Подсчёт суммы всех пар
	result := make([]int, k+1)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := (array[i] + array[j]) % mod
			// Вычисляем вклад этой пары в каждую степень p
			power := sum
			for p := 1; p <= k; p++ {
				result[p] = (result[p] + power) % mod
				power = power * sum % mod // Переход к следующей степени

			}
		}
	}

	// Вывод результатов
	for p := 1; p <= k; p++ {
		fmt.Fprintln(out, result[p])
	}
} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MOD = 998244353

var inv [301]int

func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b /= 2
	}
	return res
}

func main() {
	// Precompute inverses for 1..300
	for m := 1; m <= 300; m++ {
		inv[m] = powMod(m, MOD-2, MOD)
	}

	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	parts := strings.Split(data, " ")
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	data, _ = reader.ReadString('\n')
	data = strings.TrimSpace(data)
	aParts := strings.Split(data, " ")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(aParts[i])
		a[i] = x % MOD
	}

	// Compute S[0..k]
	S := make([]int, k+1)
	S[0] = n % MOD
	for m := 1; m <= k; m++ {
		S[m] = 0
	}
	for _, x := range a {
		xMod := x % MOD
		current := xMod
		if 1 <= k {
			S[1] = (S[1] + current) % MOD
		}
		for m := 2; m <= k; m++ {
			current = current * xMod % MOD
			S[m] = (S[m] + current) % MOD
		}
	}

	// Compute comb[p][m]
	maxP := k
	comb := make([][]int, maxP+1)
	for p := 0; p <= maxP; p++ {
		comb[p] = make([]int, p+1)
		comb[p][0] = 1
		for m := 1; m <= p; m++ {
			comb[p][m] = comb[p][m-1] * (p - m + 1) % MOD
			comb[p][m] = comb[p][m] * inv[m] % MOD
		}
	}

	// Compute pow2
	pow2 := make([]int, k+1)
	pow2[0] = 1
	for p := 1; p <= k; p++ {
		pow2[p] = pow2[p-1] * 2 % MOD
	}

	inv2 := powMod(2, MOD-2, MOD)

	// Process each p from 1 to k
	var result strings.Builder
	for p := 1; p <= k; p++ {
		sumPart := 0
		for m := 0; m <= p; m++ {
			sM := S[m]
			sPM := S[p-m]
			c := comb[p][m]
			sumPart = (sumPart + c*sM%MOD*sPM%MOD) % MOD
		}
		term := (sumPart - pow2[p]*S[p]%MOD) % MOD
		if term < 0 {
			term += MOD
		}
		res := term * inv2 % MOD
		result.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(result.String())
} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b /= 2
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Чтение входных данных
	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] %= MOD
	}

	// Предвычисление сумм степеней S[m] = sum(a_i^m) mod MOD
	S := make([]int, k+1)
	S[0] = n % MOD // a_i^0 = 1 для всех i

	for _, x := range a {
		current := x
		for m := 1; m <= k; m++ {
			S[m] = (S[m] + current) % MOD
			current = current * x % MOD
		}
	}

	// Предвычисление биномиальных коэффициентов C(p, m)
	inv := make([]int, k+1)
	for m := 1; m <= k; m++ {
		inv[m] = powMod(m, MOD-2, MOD)
	}

	comb := make([][]int, k+1)
	for p := 0; p <= k; p++ {
		comb[p] = make([]int, p+1)
		comb[p][0] = 1
		for m := 1; m <= p; m++ {
			comb[p][m] = comb[p][m-1] * (p - m + 1) % MOD
			comb[p][m] = comb[p][m] * inv[m] % MOD
		}
	}

	// Предвычисление степеней 2^p
	pow2 := make([]int, k+1)
	pow2[0] = 1
	for p := 1; p <= k; p++ {
		pow2[p] = pow2[p-1] * 2 % MOD
	}

	// Обратный элемент к 2 по модулю MOD
	inv2 := powMod(2, MOD-2, MOD)

	// Вычисление f(p) для каждого p от 1 до k
	for p := 1; p <= k; p++ {
		sumPart := 0
		for m := 0; m <= p; m++ {
			sM := S[m]
			sPM := S[p-m]
			c := comb[p][m]
			sumPart = (sumPart + c*sM%MOD*sPM%MOD) % MOD
		}
		term := (sumPart - pow2[p]*S[p]%MOD) % MOD
		if term < 0 {
			term += MOD
		}
		res := term * inv2 % MOD
		fmt.Fprintln(out, res)
	}
} */
