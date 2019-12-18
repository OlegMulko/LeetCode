package climbingstairs

import "testing"

type testPair struct {
	descriptoin string
	value       int
	average     int
}

var tests = []testPair{
	{"n = 0", 0, 1},
	{"n = 2", 2, 2},
	{"n = 3", 3, 3},
	{"n = 4", 4, 5},
	{"n = 44", 44, 1134903170},
}

func TestClimbStairs(t *testing.T) {
	for _, pair := range tests {
		v := ClimbStairs(pair.value)
		if v != pair.average {
			t.Error(
				"description", pair.descriptoin,
				"For", pair.value,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestClimbStairsMem(t *testing.T) {
	for _, pair := range tests {
		v := ClimbStairsMem(pair.value)
		if v != pair.average {
			t.Error(
				"description", pair.descriptoin,
				"For", pair.value,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			ClimbStairs(pair.value)
		}
	}
}

func BenchmarkClimbStairsMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			ClimbStairsMem(pair.value)
		}
	}
}
