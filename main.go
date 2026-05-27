package main

import (
	"fmt"
	"html/template"
	"interface/utils"
	"net/http"
)

type PageData struct {
	Input  string
	Result string
	Error  error
	Status string
}

var tmpl = template.Must(template.ParseFiles("html/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Status: "HTTP200",
	}
	tmpl.Execute(w, data)
}

func processorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	input := r.FormValue("input")
	data := PageData{
		Input: input,
	}
	result, err := utils.Multiline(input, false)
	if err != nil {
		data.Error = err
		data.Status = "HTTP400"

		w.WriteHeader(http.StatusBadRequest)

		tmpl.Execute(w, data)
		return
	}

	data.Result = "Decoded:\n" + result
	data.Status = "HTTP202"

	w.WriteHeader(http.StatusAccepted)

	tmpl.Execute(w, data)

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/processor", processorHandler)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
