package main

import (
	"errors"
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
	Status     string
}

var tmpl = template.Must(template.ParseFiles("html/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	data := PageData{Status: "200 OK", MaxPattern: 3, Mode: "decode"}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

func transformerHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	var err error
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
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
		Status:     "202 Accepted",
	}

	switch mode {
	case "decode":
		result, err = utils.Multiline(input, false, maxPattern)
	case "encode":
		result, err = utils.Multiline(input, true, maxPattern)
	default:
		err = errors.New("invalid mode")
	}

	if err != nil {
		data.Error = err
		data.Status = "400 Bad Request"
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
