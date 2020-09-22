package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Mashuri"
	names[1] = "Mansur"
	names[2] = "Muna"

	fmt.Println(len(names))

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{90, 80, 70}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
