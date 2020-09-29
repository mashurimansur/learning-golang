package main

import (
	"fmt"
)

func main() {
	//parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
	//varargs artinya datanya bisa menerima lebih dari satu input, atau sebut saja semacam array
	//jika mau menabah lebih dari 1 parameter maka parameter vargargs harus di bagian paling belakang dan paramter vargarg di sebuah function tidak bisa lebih dari 1

	total := sumAll(10, 20, 30, 40, 50, 50, 90)
	fmt.Println(total)

	//contoh jikan ingin menggunakan slice di varargs
	//untuk passing datanya maka slice harus dipecah dulu mejadi varable argument menggunakan ... (titik 3x)
	sliceNumber := []int{20, 30, 40, 50}
	totalSlice := sumAll(sliceNumber...)
	fmt.Println(totalSlice)

}

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}
