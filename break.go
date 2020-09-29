package main

import "fmt"

func main() {
	//break adalah kata kunci yang bisa digunakan dalam perulangan
	//break digunakan untuk menghentikan seluruh perulangan
	for i := 0; i < 5; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

}
