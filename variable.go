package main

import "fmt"

func main() {
	//Kalau menggunakan var, varible nya bisa di fanti nilainya
	var name string = "Hury"
	fmt.Println(name)

	name = "Mashuri Mansur"
	fmt.Println(name)

	// cara mepersingkat variable
	address := "Belawa"
	fmt.Println(address)

	//const, adalah sebuaah variable yang tidak bisa diganti nulaunya apbila sudah ditentukan sebelumnya
	const company string = "Docotel"
	fmt.Println(company)

	var (
		firstName = "Hury"
		lastName  = "Mansur"
	)

	fmt.Println(firstName, lastName)
}
