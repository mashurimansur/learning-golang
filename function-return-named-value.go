package main

import "fmt"

func main() {
	//biasanya saat kita memberitahu bahwa sebuha function mengembalikan value, kita hanya mendeklarasikan tipe data return value di function
	//namun kita juga bisa membuat variable secara langsung di tipe data return function nya
	firstName, lastName := getFullName2()
	fmt.Println(firstName, lastName)
}

func getFullName2() (firstName, lastName string) {
	firstName = "Hury"
	lastName = "Mansur"
	return
}
