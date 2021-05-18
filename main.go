package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir string = "./mock/layout"
var bootstrap *template.Template

func main() {
	var err error
	bootstrap, err = template.ParseGlob("./mock/layout/*.html")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./mock"))))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8085", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./mock/login.html"))
	t.ExecuteTemplate(w, "login", nil)
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.html")
	if err != nil {
		panic(err)
	}
	return files
}
