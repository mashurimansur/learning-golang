package main

import "fmt"

func main() {
	//saat membuat function kadang2 kita menbutuhkan data dari luar, atau kita sebut parameter
	//kita bisa menmbahkan parameter di function, bisa lebih dari satu
	//parameter tidaklah wajib

	sayHelloTo("Mashuri", "Mansur")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
