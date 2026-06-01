package main

import (
	"forum/database"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Erreur template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, "Erreur execution template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html")
}

func category(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "category.html")
}

func errorP(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "error.html")
}

func post(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "post.html")
}

func register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html")
}

func main() {
	database.InitDB()
	defer database.DB.Close()
	database.CreateTables()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html")
	})
	http.HandleFunc("/login", login)
	http.HandleFunc("/category", category)
	http.HandleFunc("/error", errorP)
	http.HandleFunc("/post", post)
	http.HandleFunc("/register", register)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
