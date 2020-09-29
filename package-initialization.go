package main

import (
	"fmt"
	"github.com/mashurimansur/learning-golang/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
