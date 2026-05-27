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

<<<<<<< HEAD
func processorHandler(w http.ResponseWriter, r *http.Request) {
=======
func processHandler(w http.ResponseWriter, r *http.Request) {
>>>>>>> 7c2f642c1f62401f9c7f2db203af8f3e26beb1fe
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	input := r.FormValue("input")
	mode := r.FormValue("mode")
	data := PageData{
		Input: input,
	}
	var result string
	var err error
	if mode == "decode" {
		result,err = utils.Multiline(input,false)
	} else if mode == "encode" {
		result, err = utils.Multiline(input,true)
	}

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

<<<<<<< HEAD
	http.HandleFunc("/processor", processorHandler)
=======
	http.HandleFunc("/process", processHandler)
>>>>>>> 7c2f642c1f62401f9c7f2db203af8f3e26beb1fe

	fmt.Println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
