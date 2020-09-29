package main

import "fmt"

func main() {
	//continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melangjutkan ke perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}
