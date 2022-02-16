package codes

import (
	"testing"
)

func TestPrime(t *testing.T) {

	testCases := []struct {
		//test   string
		input  int
		output string
	}{
		{2, "1 2 "},
		{10, "1 2 3 5 7 "},
		{50, "1 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 "},
	}

	for i, tc := range testCases {
		res := Prime(testCases[i].input)

		if res != tc.output {
			t.Errorf("testcase fails %v", i)
		}
	}
}
