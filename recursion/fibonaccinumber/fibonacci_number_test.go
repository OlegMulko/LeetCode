package fibonaccinumber

import "testing"

type testPair struct {
	descriptoin string
	value       int
	average     int
}

var tests = []testPair{
	{"N = 0", 0, 0},
	{"N = 1", 1, 1},
	{"N = 5", 5, 5},
	{"N = 8", 8, 21},
	{"N = 10", 10, 55},
	{"N = 30", 30, 832040},
	{"N = -1", -1, 0},
	{"N = 31", 31, 0},
}

func TestFib(t *testing.T) {
	for _, pair := range tests {
		v := Fib(pair.value)
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

func TestFibMem(t *testing.T) {
	for _, pair := range tests {
		v := FibMem(pair.value)
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

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			Fib(pair.value)
		}
	}
}

func BenchmarkFibMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			FibMem(pair.value)
		}
	}
}
