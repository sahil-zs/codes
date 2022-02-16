package codes

func Compute(a, b int, op string) int {
	switch op {
	case "+":
		return a + b

	case "-":
		return a - b

	case "*":
		return a * b

	case "/":
		return a / b

	default:
		return a % b
	}

}
