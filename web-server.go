package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "apa kabar")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("starting web server at :8080")
	_ = http.ListenAndServe(":8080", nil)
}
