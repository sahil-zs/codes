package codes

func PascalTriangle(n int) [][]int {
	var arr [][]int
	a := []int{1}
	arr = append(arr, a)
	if n == 1 {
		return arr
	}
	a = []int{1, 1}
	arr = append(arr, a)
	if n == 2 {
		return arr
	}
	b := []int{}
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			if j == i-1 || j == 0 {
				b = append(b, 1)
			} else {
				b = append(b, a[j-1]+a[j])
			}
		}
		arr = append(arr, b)
		a = b
		b = []int{}
	}
	return arr
}
