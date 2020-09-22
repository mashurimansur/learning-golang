package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noktp NoKTP = "111222323123"
	var marriedStatus Married = true
	fmt.Println(noktp)
	fmt.Println(marriedStatus)
}
