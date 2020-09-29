package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run flag.go -host=localhost -user=root -password=root
	host := flag.String("host", "localhost", "put your host")
	user := flag.String("user", "root", "put yout username")
	password := flag.String("password", "root", "put your password")
	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}
