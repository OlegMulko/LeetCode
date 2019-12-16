package pascalstriangle

import "testing"

type testPair struct {
	descriptoin string
	value       int
	average     [][]int
}

var tests = []testPair{
	{"numRows = -1", -1, [][]int{}},
	{"numRows = 0", 0, [][]int{}},
	{"numRows = 1", 1, [][]int{
		[]int{1},
	}},
	{"numRows = 5", 5, [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}},
}

func TestGenerate(t *testing.T) {
	error := false
	for _, pair := range tests {
		result := Generate(pair.value)
		for indexRow, row := range result {
			for index, value := range row {
				if value != pair.average[indexRow][index] {
					error = true
					break
				}
			}
		}
		if error {
			t.Error(
				"description", pair.descriptoin,
			)
		}
	}
}
