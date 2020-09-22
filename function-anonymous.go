package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
	//namun kadang ada kalanya lebih mudah membuat function secara lagnsung di variable atau parameter tanpa harus membuat function terlebih dahulu
	//hal tersebut dinamakan anonymous function, atau function tanpa nama.
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("hury", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
