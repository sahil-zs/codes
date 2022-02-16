package codes

import "testing"

func TestLeapyr(t *testing.T) {
	testCases := []struct {
		input  int
		output string
	}{
		{1997, "no"},
		{2000, "yes"},
		{1996, "yes"},
		{1900, "no"},
	}
	for i, tc := range testCases {
		res := Isleap(testCases[i].input)

		if res != tc.output {
			t.Errorf("testcase fails %v", i)
		}
	}

}

func Benchmark(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Isleap(2000)
	}
}
