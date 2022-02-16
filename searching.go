package codes

func Linear(a [10]int, elem int) bool {
	flag := false

	i := 0
	for i < len(a) {
		if a[i] == elem {
			flag = true
		}
		i++
	}
	if flag == true {
		return true
	}
	return false
}

func Binary(a [10]int, elem int) bool {
	s := 0
	e := len(a)

	for s <= e {
		var mid int = (s + e) / 2

		if a[mid] == elem {
			return true
		} else if a[mid] > elem {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return false
}
