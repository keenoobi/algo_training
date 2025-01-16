package main

import (
	"fmt"
	"math"
)

func min(A, B, C, D int) (int, int) {
	var M, N int
	MAX_M := int(math.Max(float64(B)+1, float64(A)+1))
	MAX_N := int(math.Max(float64(C)+1, float64(D)+1))

	if (A == 0 && C == 0) || (B == 0 && D == 0) {
		return 1, 1
	}

	if A == 0 || C == D { // one red and blue +1
		return 1, C + 1
	} else if B == 0 { // one blue and red + 1
		return 1, D + 1
	} else if C == 0 || A == B || (A == B && A == C && A == D) { // blue + 1 and one red
		return A + 1, 1
	} else if D == 0 { // red + 1 and one blue
		return B + 1, 1
	}

	SUM1 := MAX_M + 1
	SUM2 := MAX_N + 1

	SUM3 := A + C + 2
	SUM4 := B + D + 2

	MIN1 := int(math.Min(float64(SUM1), float64(SUM2)))
	MIN2 := int(math.Min(float64(SUM3), float64(SUM4)))
	RESULT := int(math.Min(float64(MIN1), float64(MIN2)))

	switch RESULT {
	case SUM1:
		M, N = MAX_M, 1
	case SUM2:
		M, N = 1, MAX_N
	case SUM3:
		M, N = A+1, C+1
	case SUM4:
		M, N = B+1, D+1
	}
	return M, N
}

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	M, N := min(A, B, C, D)

	fmt.Print(M, N)

}
