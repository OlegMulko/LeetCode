package pascalstriangle2

import "testing"

type testPair struct {
	descriptoin string
	value       int
	average     []int
}

var tests = []testPair{
	{"rowIndex = -1", -2, []int{}},
	{"rowIndex = 0", 0, []int{1}},
	{"rowIndex = 1", 1, []int{1, 1}},
	{"rowIndex = 3", 3, []int{1, 3, 3, 1}},
}

func TestGetRow(t *testing.T) {
	error := false
	for _, pair := range tests {
		result := GetRow(pair.value)
		for index, value := range result {
			if value != pair.average[index] {
				error = true
				break
			}
		}
		if error {
			t.Error(
				"description", pair.descriptoin,
			)
		}
	}
}
