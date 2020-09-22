package main

import "fmt"

/**
interface adalah tipe data abstract, di tidak memiliiki implementasi langsung
sebuah interface berisikan definisi-definisi method
biasanya interface digunakan sebagai kontrak
*/
type HashName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHelllo(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

/**
setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
sehingga kita tidak perlu mengimplementasikan interface secara manual
hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara explisit akan menggunakan interace mana
*/
func main() {
	var hury Person
	hury.Name = "Mashuri"
	sayHelllo(hury)

	var cat Animal
	cat.Name = "Kucing"
	sayHelllo(cat)
}
