package codes

//Determine if a triangle is equilateral, isosceles, or scalene.
//
//An equilateral triangle has all three sides the same length.
//
//An isosceles triangle has at least two sides the same length. (It is sometimes specified as having exactly two sides the same length, but for the purposes of this exercise we'll say at least two.)
//
//A scalene triangle has all sides of different lengths.

func Determine(a, b, c int) string {
	if a == b && b == c && c == a {
		return "equilateral"
	} else if a == b || b == c || c == a {
		return "isosceles"
	} else {
		return "scalene"
	}
}
