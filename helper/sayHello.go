package helper

import "fmt"

var version = "1.0.0"
var Application = "Aplikasi Belajar"

func SayHello(name string) string {
	return "Hello" + name
}

func sayGoodBye(name string) {
	fmt.Println("Bye Bye", name)
}
