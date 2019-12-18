package climbingstairs

var climbStairsKСache = map[int]int{}

// ClimbStairs ...
func ClimbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return ClimbStairs(n-2) + ClimbStairs(n-1)
}

// ClimbStairsMem ...
func ClimbStairsMem(n int) int {
	var result int

	if KeyVelue, KeyExist := climbStairsKСache[n]; KeyExist {
		return KeyVelue
	}
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	result = ClimbStairsMem(n-2) + ClimbStairsMem(n-1)
	climbStairsKСache[n] = result
	return result
}
