package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	duration := time.Since(start)

	fmt.Println("time elapsed in second", duration.Seconds())
	fmt.Println("time elapsed in minutes", duration.Minutes())
	fmt.Println("time elapsed in hour", duration.Hours())
}
