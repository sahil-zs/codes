package codes

func Passbyvaluee(a, b int) {
	a = a + 1

}

func Passbyreferencee(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}
