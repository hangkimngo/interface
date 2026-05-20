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
	Status int
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Status: 200,
	}
	tmpl.Execute(w, data)
}

func decoderHandler(w http.ResponseWriter, r *http.Request) {
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
		data.Status = 400

		w.WriteHeader(http.StatusBadRequest)

		tmpl.Execute(w, data)
		return
	}

	data.Result = "Decoded:\n" + result
	data.Status = 202

	w.WriteHeader(http.StatusAccepted)

	tmpl.Execute(w, data)

}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/decoder", decoderHandler)

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
