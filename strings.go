package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Mashuri Mansur"
	fmt.Println(strings.Contains(name, "huri"))                  // return boolean
	fmt.Println(strings.Split(name, " "))                        // return []string
	fmt.Println(strings.ToLower(name))                           //return string
	fmt.Println(strings.ToUpper(name))                           // return string
	fmt.Println(strings.Trim("         hai              ", " ")) // return string
	fmt.Println(strings.ReplaceAll(name, "a", "i"))              // return string
}
