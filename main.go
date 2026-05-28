package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/processor", processorHandler)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
