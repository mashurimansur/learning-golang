package main

import "fmt"

func main() {
	//closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
	//harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
	//saat kita salah menggunakan closure maka kita sudah mengganti data yang sebenarnya tidak perlu diubah

	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
