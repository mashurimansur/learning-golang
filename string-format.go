package main

import "fmt"

type pelajar struct {
	Name        string
	Height      float64
	Age         int
	IsGraduated bool
	Hobbies     []string
}

func main() {
	var data = pelajar{
		Name:        "Mashuri",
		Height:      175,
		Age:         23,
		IsGraduated: true,
		Hobbies:     []string{"eating", "sleeping"},
	}

	// mengubah data numerik menjadi string numberik berbasis biner
	fmt.Printf("%b\n", data.Age)
	// 11010

	//memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicodenya
	fmt.Printf("%c\n", 1400)
	// ո
	fmt.Printf("%c\n", 1235)
	// ӓ

	//meformat data numerik menjadi bentuk string numerik berbasi 10 (basis bilangan yang kita gunakan)
	fmt.Printf("%d \n", data.Age)
	// 26

	//Digunakan untuk memformat data numerik desimal ke dalam bentuk notasi numerik standar Scientific notation.
	fmt.Printf("%e\n", data.Height)
	// 1.825000e+02
	fmt.Printf("%E\n", data.Height)
	// 1.825000E+02

	/**
	%F adalah alias dari %f. Keduanya memiliki fungsi yang sama.
	Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.
	*/
	fmt.Printf("%f\n", data.Height)
	// 182.500000
	fmt.Printf("%.9f\n", data.Height)
	// 182.500000000
	fmt.Printf("%.2f\n", data.Height)
	// 182.50
	fmt.Printf("%.f\n", data.Height)
	// 182

	fmt.Printf("%+v\n", data)
	// {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}
}
