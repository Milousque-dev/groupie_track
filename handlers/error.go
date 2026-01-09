package handlers

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Code    int
	Message string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.WriteHeader(code)

	data := ErrorData{
		Code:    code,
		Message: message,
	}

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, message, code)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, message, code)
		return
	}
}
