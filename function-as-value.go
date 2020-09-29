package main

import "fmt"

func main() {
	//function adalah first class citizen
	//function juga merupakan tipe data, dan bisa disimpan dalam variabel
	sayGoodBye := getGoodBye
	result := sayGoodBye("Hury")
	fmt.Println(result)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
