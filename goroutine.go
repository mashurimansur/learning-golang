package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "Hallo")
	print(5, "Apa Kabar")

	var input string
	_, _ = fmt.Scanln(&input)
	fmt.Println(input)
}
