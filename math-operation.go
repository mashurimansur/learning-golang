package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	//Augmented Assigments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	//Unary Operator
	i++ // i = i + 1
	fmt.Println(i)
}
