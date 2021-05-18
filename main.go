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

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./mock/static"))))
	http.HandleFunc("/login", login)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/users", users)
	http.HandleFunc("/usersadd", usersadd)
	http.HandleFunc("/userprofile", userprofile)
	http.ListenAndServe(":8085", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./mock/login.html"))
	err := t.ExecuteTemplate(w, "login", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/dashboard.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/users.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func usersadd(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/users_add.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userprofile(w http.ResponseWriter, r *http.Request) {
	files := append(layoutFiles(), "./mock/user.html")
	t := template.Must(template.ParseFiles(files...))
	err := t.ExecuteTemplate(w, "bootstrap", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.html")
	if err != nil {
		panic(err)
	}
	return files
}
