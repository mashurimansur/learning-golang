package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	// unutuk menambahkan data paling belakang
	data.PushBack("Huri")
	data.PushBack("Mansur")
	data.PushBack("Muna")

	//untuk menambahkan data paling depan
	data.PushFront("Jumardi")

	// untuk mengambil data paling depan
	fmt.Println(data.Front().Value)
	//unutuk mengambil data paling belakang
	fmt.Println(data.Back().Value)

	// menampilkan data dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// menampilkan data deri belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
