package main

import "fmt" 

func main() {
	// Manipulasi string
	fmt.Println("Belajar Golang")

	//Menggabungkan string
	fmt.Println("Belajar" + " " + "Golang")

	// mengetahui panjang string
	fmt.Println(len("Belajar"))

	// mengambil huruf pertama, namun outputnya adalah berupa byte, harus di konversi terlebih dahulu jadi string
	fmt.Println("Belajar"[0])

	// mengambil beberapa kata dari string, 1 artinya index ke 1 dan 4 adalah posisi huruf yang diinginkan
	fmt.Println("Belajar"[1:4])

	//kalau kita ingin mengabil string dari posisi 2 dan diberi tanda ":" lalu kosong setelahnya artinya mengambil string dari index 2 dan setrusnya
	fmt.Println("Belajar"[2:])

	//sebelum ":" itu kosong attinya mengambil string dari awal sampai ke urutan 4
	fmt.Println("Belajar"[:4])
}