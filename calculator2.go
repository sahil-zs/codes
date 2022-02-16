package codes

var (
	Add = func(x, y int) int {
		return x + y
	}

	Diff = func(x, y int) int {
		return x - y
	}

	Mul = func(x, y int) int {
		return x * y
	}

	Div = func(x, y int) int {
		return x / y
	}
)

func Calculate(x, y int, op func(x, y int) int) int {
	res := op(x, y)
	return res
}
