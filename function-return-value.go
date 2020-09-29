package main

import "fmt"

func main() {
	//function bisa mengembalikan data
	//untuk meberitahu bahwa function mengembalikan data, kita harus menuliskan tipde data kembalian dari function tersebut
	//jika function tersebut kita deklarasikan dengan tipe data kembalian
	result := getHello("Hury")
	fmt.Println(result)
}

func getHello(name string) string {
	return "Hello " + name
}
