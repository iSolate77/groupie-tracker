package handlers

import (
    "net/http"
    "html/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
    tmpl.Execute(w, nil)
}
