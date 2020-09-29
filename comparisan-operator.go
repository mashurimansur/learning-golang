package main

import "fmt"

func main() {
	var name1 = "Hury"
	var name2 = "Hury"

	result := name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 100

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
