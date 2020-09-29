package main

import "fmt"

import "strconv"

func main() {
	var x int32 = 10
	var y int64 = int64(x)
	fmt.Println(y)

	var z float64 = float64(y)
	fmt.Println(z)

	var nilai string = "100"
	nilaiInt, _ := strconv.Atoi(nilai)
	fmt.Println(nilaiInt)

	nilaiString := strconv.Itoa(nilaiInt)
	fmt.Println(nilaiString)

	var name = "Hury"
	var h = name[0]
	var hString = string(h)
	fmt.Println(hString)
}
