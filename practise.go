package codes

import "fmt"

const (
	// Big the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Small Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

const pi = 3.14

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func try() {
	fmt.Println(needInt(Small))
	//fmt.Println(needInt())
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(pi)
}
