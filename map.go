package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Hury",
		"address": "Makassar",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Saya"
	book["upss"] = "Uppss"

	fmt.Println(book)

	delete(book, "upss")
	fmt.Println(book)
}
