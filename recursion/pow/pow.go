package pow

// MyPow ...
func MyPow(x float64, n int) float64 {
	var result float64
	isNegative := false

	if n < 0 {
		isNegative = true
		n = n * -1
	}
	switch {
	case n == 0:
		return 1
	case n == 1:
		result = x
	default:
		result = x * MyPowOpt(x, n-1)
	}
	if isNegative {
		return 1 / result
	}
	return result
}

// MyPowOpt ...
func MyPowOpt(x float64, n int) float64 {
	var result float64
	isNegative := false

	if n < 0 {
		isNegative = true
		n = n * -1
	}
	switch {
	case n == 0:
		return 1
	case n == 1:
		result = x
	case n%2 == 0:
		result = MyPowOpt(x*x, n/2)
	default:
		result = x * MyPowOpt(x, n-1)
	}
	if isNegative {
		return 1 / result
	}
	return result
}
