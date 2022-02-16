package codes

import "testing"

func TestIdentify(t *testing.T) {
	var testCases = []struct {
		side1  int
		side2  int
		side3  int
		output string
	}{
		{3, 3, 3, "equilateral"},
		{2, 3, 2, "isosceles"},
		{1, 2, 3, "scalene"},
	}
	for i, tc := range testCases {
		res := Determine(testCases[i].side1, testCases[i].side2, testCases[i].side3)

		if res != tc.output {
			t.Errorf("testcase fails %v", i)
		}
	}

}

func BenchmarkIdentify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Determine(3, 4, 5)
	}
}
