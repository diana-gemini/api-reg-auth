package render

import (
	"fmt"
	"html/template"
	"net/http"
)

type WebPage struct {
	IsLoggedin bool
	Errtext    string
}

func Render(w http.ResponseWriter, temp string, data WebPage) {

	templ, err := template.ParseFiles(temp, "ui/html/layout.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
