package main

import "fmt"

func main() {
	var name = "Hury"

	if name == "Hury" {
		fmt.Println("Hello Hury")
	} else if name == "Mansur" {
		fmt.Println("Hallo, Pak Mansur")
	} else {
		fmt.Println("hai boleh kenalan nggak?")
	}

	//if short statement
	//contoh mendeklarasikan variabel di dalam if
	//sehingga lebih simpel
	//mmungkin ini cuman ada di bahasa golang
	if lenght := len(name); lenght > 5 {
		fmt.Println("lebih dari 5 huruf")
	} else {
		fmt.Println("Namanya terlalu pendek")
	}
}
