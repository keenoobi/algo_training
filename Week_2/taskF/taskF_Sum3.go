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

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fstr := strings.Fields(input)
	N, _ := strconv.Atoi(fstr[0])

	scndstr, _ := reader.ReadString('\n')
	scndstr = strings.TrimSpace(scndstr)
	elements := strings.Fields(scndstr)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i], _ = strconv.Atoi(elements[i])
	}

	// Вычисляем суммы S1, S2, S3
	MOD := 1000000007
	S1, S2, S3 := 0, 0, 0
	for _, num := range array {
		S1 = (S1 + num) % MOD
		S2 = (S2 + num*num) % MOD
		S3 = (S3 + num*num*num) % MOD
	}

	// Вычисляем результат по формуле
	result := (S1*S1%MOD*S1%MOD + 2*S3%MOD - 3*S1%MOD*S2%MOD + MOD) % MOD
	result = (result * 166666668) % MOD // 166666668 = 1/6 mod 1000000007

	/* // Массив для хранения сумм всех элементов после j
	sumK := make([]int, N)
	// Массив для хранения сумм всех произведений array[j] * sumK[j] для всех j > i
	sumJ := make([]int, N)

	// Заполняем массивы sumK и sumJ
	for i := N - 1; i >= 0; i-- {
		sumK[i] = array[i]
		if i < N-1 {
			sumK[i] = (sumK[i] + sumK[i+1]) % MOD
		}
		if i < N-2 {
			sumJ[i] = (array[i+1] * sumK[i+2]) % MOD
			if i < N-3 {
				sumJ[i] = (sumJ[i] + sumJ[i+1]) % MOD
			}
		}
	}

	// Вычисляем результат
	result := 0
	for i := 0; i < N-2; i++ {
		result = (result + array[i]*sumJ[i]) % MOD
	} */

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, result)
	writer.Flush()
}
