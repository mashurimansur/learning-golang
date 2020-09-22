package main

import (
	"errors"
	"fmt"
)

/**
golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error
untuk membuat error, kita tidak perlu manual
golang sudah menyediakan library untuk membuat error, yaitu package errors
*/
func Pembagi(nilai int, pembagai int) (int, error) {
	if pembagai == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan 0")
	} else {
		return nilai / pembagai, nil
	}
}

func main() {
	result, err := Pembagi(10, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
