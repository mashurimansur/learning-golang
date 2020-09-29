package main

import "fmt"

/**
struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
struct biasanya representasi data dalam program aplikasi yang kita buat
data di struct disimpan dalam field
sederhananya struct adalah kumpulan field
programmer golang biasanya menggunakan uppercase dalam penamaan struct
*/

type Customer struct {
	Name, Address string
	Age           int
}

//membuat data struct
/**
struct adalah template data atay protorype data
struct tidak bisa langsung digunakan
namun kita bisa membuat data/object dari struct yang telah kita buat
*/

func main() {
	var hury Customer
	hury.Name = "Mashuri"
	hury.Address = "Makassar"
	hury.Age = 23

	fmt.Println(hury)
	fmt.Println(hury.Name)
	fmt.Println(hury.Address)
	fmt.Println(hury.Age)

	/**
	struct literals
	sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
	*/
	//contoh 1
	juma := Customer{
		Name:    "Jumardi",
		Address: "Belawa",
		Age:     20,
	}
	fmt.Println(juma)

	//contoh 2, contoh ini tidak disarankan karena ketika terjadi perubahan posisi data atau penambahan data di struct makan contoh ini akan error
	mansur := Customer{"Mansur", "Belawa", 40}
	fmt.Println(mansur)

	/**
	Struct Method
	struct adalah tipe data seperti tipe data lainnya, di bisa digunakan sebagai parameter unutk function
	namun jika kita ingin menmbahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
	method adalah function
	*/
	muna := Customer{Name: "Muna", Address: "Belawa"}
	muna.seyHi()
	muna.sayHuuu()
}

//contoh struct method
func (customer Customer) seyHi() {
	fmt.Println("Hello, My name is", customer.Name)
}

// contoh struct method
func (customer Customer) sayHuuu() {
	fmt.Println("huuuuuuuuuuuuuuuuuuuuuu from", customer.Address)
}
