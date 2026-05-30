package main

import (
	"html/template"
	"interface/utils"
	"net/http"
	"strconv"
)

type PageData struct {
	Mode       string
	MaxPattern int
	Input      string
	Result     string
	Error      error
}

var tmpl = template.Must(template.ParseFiles("html/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Mode:       "decode",
		MaxPattern: 3,
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

func processorHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	var err error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	input := r.FormValue("input")
	mode := r.FormValue("mode")
	maxPattern := 3
	if v := r.FormValue("max-pattern"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			maxPattern = n
		}
	}

	data := PageData{
		Mode:       mode,
		Input:      input,
		MaxPattern: maxPattern,
	}

	if mode == "decode" {
		result, err = utils.Multiline(input, false, maxPattern)
	} else {
		result, err = utils.Multiline(input, true, maxPattern)
	}

	if err != nil {
		data.Error = err
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, data)
		return
	}

	if mode == "encode" {
		data.Result = result
	} else {
		data.Result = result
	}
	w.WriteHeader(http.StatusAccepted)
	tmpl.Execute(w, data)
}
