package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	// go run os.go mashuri mansur
	//fmt.Println(args[1])
	//fmt.Println(args[2])

	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname")
	} else {
		fmt.Println("Hostname:", name)
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}
