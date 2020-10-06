package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak bisa kosong")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("application running perfectly")
	}
}

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	_, _ = fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hallo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
