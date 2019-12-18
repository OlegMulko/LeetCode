package fibonaccinumber

var climbStairsKFib = map[int]int{}

// Fib ...
func Fib(N int) int {
	if N <= 0 || N > 30 {
		return 0
	}
	if N > 2 {
		return Fib(N-1) + Fib(N-2)
	}
	return 1
}

// FibMem ...
func FibMem(N int) int {
	var result int

	if KeyValue, KeyExist := climbStairsKFib[N]; KeyExist {
		return KeyValue
	}
	if N <= 0 || N > 30 {
		return 0
	}
	if N > 2 {
		result = FibMem(N-1) + FibMem(N-2)
		climbStairsKFib[N] = result
		return result
	}
	return 1
}
