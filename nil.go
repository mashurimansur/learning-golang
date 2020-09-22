package main

import "fmt"

/**
biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
berbeda degan golang, di golang kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
namun di golang ada data nil, yaitu data kosong
nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
*/
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	name := NewMap("Hury")
	fmt.Println(name["name"])
}
