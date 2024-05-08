package funkcje

import "math/big"

var LiczbaWywołańFib = map[int]int{}
var CałkowitaLiczbaWywołańFib int = 0

func Silnia(n int) *big.Int {
	return big.NewInt(0).MulRange(1, int64(n))
}

func Fibonacci(n int) int {
	LiczbaWywołańFib[n] += 1
	CałkowitaLiczbaWywołańFib += 1
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
