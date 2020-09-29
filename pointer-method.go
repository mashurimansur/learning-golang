package main

import "fmt"

type Man struct {
	Name string
}

/**
walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
sangan direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu duplikasi ketikan memanggil method
*/
func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}
func main() {
	hury := Man{"Hury"}
	hury.Merried()
	fmt.Println(hury.Name)
}
