package codes

import (
	"strconv"
)

func printPrime(value1 int, s string) string {
	var v = 2
	flag := true
	//var s string=""
	for v <= value1/2 {
		if value1%v == 0 {
			flag = false
		}
		v++
	}
	if flag == true {
		s += strconv.Itoa(value1)
		s += " "
	}
	return s
}

func Prime(val int) string {
	s := ""
	cnt := 1
	for cnt <= val {

		s = printPrime(cnt, s)
		cnt++
	}
	return s
}
