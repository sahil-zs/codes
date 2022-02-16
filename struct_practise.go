package codes

import "fmt"

type details struct {
	name string
	sal  int
	age  int
}

type employee struct {
	id   int
	info details
}

func main1() {

	result := employee{

		info: details{"Sahil", 10000000, 25},
	}

	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}
