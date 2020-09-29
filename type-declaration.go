package main

import "fmt"

func main() {
	type Married bool
	type NoKTP string

	var noktp NoKTP = "111222323123"
	var marriedStatus Married = true
	fmt.Println(noktp)
	fmt.Println(marriedStatus)
}
