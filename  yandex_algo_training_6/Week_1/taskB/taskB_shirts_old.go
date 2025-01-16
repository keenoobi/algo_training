package main

import (
	"fmt"
	"math"
)

// func min(A, B, C, D int) (int, int) {
// 	// minM := 1000000001
// 	// minN := 1000000001
// 	var M, N int

// 	if (A == 0 && C == 0) || (B == 0 && D == 0) { // both zeros
// 		return 1, 1
// 	}
// 	if A == 0 || B == 0 { // shirts = zero
// 		M = 1
// 		N = int(math.Max(float64(C), float64(D))) + 1
// 		return M, N
// 	} else if C == 0 || D == 0 { // socks = zero
// 		M = int(math.Max(float64(A), float64(B))) + 1
// 		N = 1
// 		return M, N
// 	}

// 	if A == B { // shirts are equal
// 		if A+B > C+D {
// 			M = 1
// 			N = A + 1
// 			return M, N
// 		} else {
// 			M = A + 1
// 			N = 1
// 			return M, N
// 		}
// 	} else if C == D { // socks are equal
// 		if A+B > C+D {
// 			M = 1
// 			N = C + 1
// 			return M, N
// 		} else {
// 			M = C + 1
// 			N = 1
// 			return M, N
// 		}
// 	}
// 	/*

// 	 */

// 	if A+C > C+D
// 	fmt.Println(A+C, B+D)
// 	fmt.Println(A+B, C+D)
// 	if A+C > B+D { // blue > red
// 		if A+B > C+D {

// 		}
// 		M = int(math.Min(float64(A), float64(B))) + 1
// 		N = int(math.Min(float64(C), float64(D))) + 1
// 		return M, N
// 	} else { // red > blue
// 		if A+B > C+D {
// 			M = int(math.Max(float64(A), float64(B))) + 1
// 			N = 1
// 			return M, N
// 		} else {
// 			M = int(math.Min(float64(A), float64(B))) + 1
// 			N = int(math.Min(float64(C), float64(D))) + 1
// 			return M, N
// 		}
// 	}

// 	if A+B < C+D { // socks > shirts
// 		if A+C > B+D { // blue > red
// 			M = 1
// 			N = int(math.Max(float64(C), float64(D))) + 1
// 			return M, N
// 		} else {
// 			M = int(math.Min(float64(A), float64(B))) + 1
// 			N = int(math.Min(float64(C), float64(D))) + 1
// 			return M, N
// 		}
// 	} else if A+B > C+D { // shirts > socks
// 		if A+C > B+D { // blue > red
// 			M = 1
// 			N = int(math.Max(float64(C), float64(D))) + 1
// 			return M, N
// 		} else {
// 			M = int(math.Min(float64(A), float64(B))) + 1
// 			N = int(math.Min(float64(C), float64(D))) + 1
// 			return M, N
// 		}
// 	}

// 	// M = int(math.Min(float64(A), float64(B))) + 1
// 	// N = int(math.Min(float64(C), float64(D))) + 1
// 	// return M, N

// 	// if A+B < C+D {
// 	// 	if A+C > B+D {
// 	// 		M = 1
// 	// 		N = int(math.Max(float64(B), float64(D))) + 1
// 	// 		return M, N
// 	// 	} else {
// 	// 		M = int(math.Max(float64(A), float64(B))) + 1
// 	// 		N = 1
// 	// 		return M, N
// 	// 	}
// 	// } else if A+B > C+D {
// 	// 	M = int(math.Max(float64(A), float64(B))) + 1
// 	// 	N = 1
// 	// 	return M, N
// 	// } else if A == B && A == C && A == D {
// 	// 	M = A + 1
// 	// 	N = 1
// 	// 	return M, N
// 	// }

// 	// if A > 0 && C > 0 && (B != 0 && D != 0) {
// 	// 	M := A + 1
// 	// 	N := C + 1
// 	// 	if M+N < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}

// 	// }

// 	// if B > 0 && D > 0 && (A != 0 && C != 0) {
// 	// 	M := B + 1
// 	// 	N := D + 1
// 	// 	if M+N < minM+minN {
// 	// 		minM, minN = M, N

// 	// 	}
// 	// }

// 	// if A == 0 && B > 0 {
// 	// 	M := 1
// 	// 	N := int(math.Max(float64(C), float64(D))) + 1
// 	// 	if M+N < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if B == 0 && A > 0 {
// 	// 	M := 1
// 	// 	N := int(math.Max(float64(C), float64(D))) + 1
// 	// 	if M+N < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if C == 0 && D > 0 {
// 	// 	M := int(math.Max(float64(A), float64(B))) + 1
// 	// 	N := 1
// 	// 	if M+N < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if D == 0 && C > 0 {
// 	// 	M := int(math.Max(float64(A), float64(B))) + 1
// 	// 	N := 1
// 	// 	if M+N < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if A > 0 {
// 	// 	M := 1
// 	// 	N := A + 1
// 	// 	if M < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if B > 0 {
// 	// 	M := 1
// 	// 	N := B + 1
// 	// 	if M < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if C > 0 {
// 	// 	M := 1
// 	// 	N := C + 1
// 	// 	if M < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	// if D > 0 {
// 	// 	M := 1
// 	// 	N := D + 1
// 	// 	if M < minM+minN {
// 	// 		minM, minN = M, N
// 	// 	}
// 	// }

// 	return M, N
// }

func min11(A, B, C, D int) (int, int) {
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

	// if MAX_M+1 < A+C+2 {
	// 	return 1, B + 1
	// } else if MAX_N+1 < B+D+2 {
	// 	return A + 1, 1
	// }
	// M2 := int(math.Max(float64(B)+1, float64(A)+1))
	// N2 := int(math.Min(float64(C)+1, float64(D)+1))
	// S2 := M2 + N2
	// fmt.Println(S2)

	// M3 := int(math.Min(float64(A)+1, float64(B)+1))
	// N3 := int(math.Max(float64(D)+1, float64(C)+1))
	// S3 := M3 + N3
	// fmt.Println(S3)

	// M4 := int(math.Max(float64(A)+1, float64(B)+1))
	// N4 := int(math.Max(float64(C)+1, float64(D)+1))
	// S4 := M4 + N4
	// fmt.Println(S4)

	// if S1 <= S2 && S1 <= S3 && S1 <= S4 {
	// 	fmt.Println("1")
	// 	return M1, N1
	// } else if S2 <= S1 && S2 <= S3 && S2 <= S4 {
	// 	fmt.Println("2")
	// 	return M2, N2
	// } else if S3 <= S1 && S3 <= S2 && S3 <= S4 {
	// 	fmt.Println("3")
	// 	return M3, N3
	// } else {
	// 	fmt.Println("4")
	// 	return M4, N4
	// }
}

func test() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	M, N := min(A, B, C, D)

	fmt.Print(M, N)

}
