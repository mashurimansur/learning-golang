package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	//for dengan statement
	for hitung := 1; hitung <= 10; hitung++ {
		fmt.Println("perhitungan ke", hitung)
	}

	//for range
	//for range sangat bagus digunakan pada slice array dan map
	slice := []string{"Mashuri", "Mansur", "Muna", "Jumardi", "Halimah"}
	for key, value := range slice {
		fmt.Println("index", key, "=", value)
	}

	person := make(map[string]string)
	person["job"] = "Programmer"
	person["name"] = "Hury"

	for k, v := range person {
		fmt.Println(k, "=", v)
	}
}
