package main

import "fmt"

/**
Interface Kosong
golang bukanlah bahasa pemrograman yang berorientasi objek
biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
contoh di java ada java.lang.Object
untuk menangani kasus seperti ini, di golang kita bisa menggunakan interface kosong
interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(1)
	fmt.Println(data)
}
