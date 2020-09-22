package main

import "fmt"

/**
secara default di golang semua variable itu dipassing by value, bukan by reference
artinya, jika kita mengirim sebuah variable ke dalam function, method atau variabel lain, sebenarnya yang dikirim adalah duplikasi valuenya
*/
type Address struct {
	City     string
	Province string
	Country  string
}

/**
pointer di function
saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah
hal ini membuat variable menjadi akan, karena tidak akan bisa berubah
namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
untuk melakukan ini, kita juga bisa menggunakan pointer di function
untuk menjadikan sebauh parmeter sebagai pointer, kita bisan menggunakan operator "*" di parameternya
*/
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

/**
pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
sederhanya, dengan kemampuan pointer, kita bisa membuat pass by reference

operator yg digunakan dalam pointer
1. Operator &
untuk membuat sebuah variabel dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator "&" diikuti dengan nama variable nya.
2. Operator *
saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
semua variabel yang mengacu ke data yang sama tidak akan berubah
jika kita ingin mengubah seluruh variabel yang mengacu ke data tersebut, kita bisa menggunakan operator "*"
3. Function new()
sebelumnya untuk membuat pointer dengan menggunakan operator "&"
golang juga memiliki function new yang bisa digunakan untuk membuat pointer
namun function new hanya mengembalikan pointer data kosong, artinya tidak ada data awal
*/
func main() {
	address1 := Address{"Belawa", "Sulsel", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"
	*address2 = Address{"Jakarta", "DKI", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Belawa"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Sengkang",
		Province: "Sulsel",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}

/**
gopath adalah environment variable yang berisikan lokasi tempat kita menyimpan project dan library golang
gopath wajib dibuat ketika kita mulai membuat aplikasi lebih dari satu file atau butuh menggunakan library dari luar
*/
