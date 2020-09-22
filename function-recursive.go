package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	//recursive function adalah function yang memanggil function dirinya sendiri
	//kadang dalam pekerjaan, kita sering menemui kasus diaman menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recirsive function
	//contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial

	//contoh factorial menggunakan looping
	loop := factorialLoop(5)
	fmt.Println(loop)

	//contoh factorial menggunakan recursive function
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
