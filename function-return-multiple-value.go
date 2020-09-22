package main

import "fmt"

func main() {
	//function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
	//untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return valuenya di function
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

func getFullName() (string, string) {
	return "Mashuri", "Mansur"
}
