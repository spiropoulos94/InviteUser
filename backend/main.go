package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hey!")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Heeeeey"))
	})


	http.ListenAndServe(":8080", nil)
}