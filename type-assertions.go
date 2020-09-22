package main

import "fmt"

func randomString() interface{} {
	return 1231231
}

func randomBool() interface{} {
	return true
}

/**
type assertion merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/
func main() {
	var result interface{} = randomString()
	//resultString := result.(string)
	//fmt.Println(resultString)
	//
	//var result2 interface{} = randomBool()
	//var resultBool bool = result2.(bool)
	//fmt.Println(resultBool)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("unknown type")
	}

}
