package pow

import "testing"

type values struct {
	x float64
	n int
}

type testPair struct {
	descriptoin string
	value       values
	average     float64
}

var tests = []testPair{
	{"Pow = -2", values{x: 2.0, n: -2}, 0.25},
	{"Pow = 0", values{x: 2.0, n: 0}, 1.0},
	{"Pow = 1", values{x: 2.0, n: 1}, 2.0},
	{"Pow = 3", values{x: 2.0, n: 3}, 8.0},
	{"Pow = 8", values{x: 2.0, n: 8}, 256.0},
	{"Pow = 16", values{x: 2.0, n: 32}, 4.294967296e+09},
}

func TestMyPow(t *testing.T) {
	error := false
	for _, pair := range tests {
		result := MyPow(pair.value.x, pair.value.n)
		if result != pair.average {
			error = true
		}
		if error {
			t.Error(
				"description", pair.descriptoin,
				"for x = ", pair.value.x,
				"for n = ", pair.value.n,
				"average = ", pair.average,
				"v = ", result,
			)
		}
	}
}

func TestMyPowOpt(t *testing.T) {
	error := false
	for _, pair := range tests {
		result := MyPowOpt(pair.value.x, pair.value.n)
		if result != pair.average {
			error = true
		}
		if error {
			t.Error(
				"description", pair.descriptoin,
				"for x = ", pair.value.x,
				"for n = ", pair.value.n,
				"average = ", pair.average,
				"v = ", result,
			)
		}
	}
}

func BenchmarkMyPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			MyPow(pair.value.x, pair.value.n)
		}
	}
}

func BenchmarkMyPowOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range tests {
			MyPowOpt(pair.value.x, pair.value.n)
		}
	}
}
