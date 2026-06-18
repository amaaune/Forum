package main

import (
	"forum/database"
	"forum/handlers"
	"forum/middleware"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles(
		"templates/"+tmpl,
		"templates/poly/header.html",
		"templates/poly/footer.html",
		"templates/poly/modal.html",
	)
	if err != nil {
		http.Error(w, "Erreur template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// CORRECTION : envoyer les données au template
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Erreur execution template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteSession(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	database.InitDB()
	defer database.DB.Close()

	database.CreateTables()
	database.SeedDatabase()

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./static")),
		),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.IndexHandler(w, r, renderTemplate)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "login.html", nil)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "register.html", nil)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorHandler(w, r, renderTemplate)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		handlers.SinglePostHandler(w, r, renderTemplate)
	})

	http.HandleFunc("/post/create", handlers.CreatePostHandler)

	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)

	http.HandleFunc("/post/vote", handlers.VoteHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}