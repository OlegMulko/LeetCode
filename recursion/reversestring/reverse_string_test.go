package reversestring

import "testing"

type testPair struct {
	descriptoin string
	value       []byte
	average     []byte
}

var tests = []testPair{
	{"test1", []byte{'a'}, []byte{'a'}},
	{"test2", []byte{'m', 'e'}, []byte{'e', 'm'}},
	{"test3", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	{"test4", []byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
}

func TestReverseString(t *testing.T) {
	error := false
	for _, pair := range tests {
		ReverseString(pair.value)
		for index, value := range pair.value {
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
