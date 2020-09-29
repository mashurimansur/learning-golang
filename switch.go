package main

import "fmt"

func main() {
	name := "Mansur"

	switch name {
	case "Hury":
		fmt.Println("hai Hury")
	case "Mansur":
		fmt.Println("Hai Pake Mansur")
	default:
		fmt.Println("Hai boleh kenalan nggak?")
	}

	//swicth dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//switch tanpa kondisi
	switch {
	case len(name) > 10:
		fmt.Println("nama terlalu panjang")
	case len(name) > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("Ok sudah pas")
	}
}
