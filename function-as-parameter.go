package main

import "fmt"

//function type decalaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//function tidak hanya bisa kita simpan di dalam variable sebagai value
	//namun juga bisa kita gunakan sebagai parameter untuk function lain
	sayHelloWithFilter("Hury", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
