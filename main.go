package main

import (
	"forum/database"
	"forum/handlers"
	"forum/middleware"
	"forum/models"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data []models.Post) {
	t, err := template.ParseFiles("templates/"+tmpl, "templates/poly/header.html",
		"templates/poly/footer.html")
	if err != nil {
		http.Error(w, "Erreur template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Erreur execution template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", nil)
}

func errorP(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "error.html", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html", nil)
}

func post(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "post.html", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteSession(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	database.InitDB()
	defer database.DB.Close()
	database.CreateTables()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts, err := handlers.GetPosts()
		if err != nil {
			return
		}
		renderTemplate(w, "index.html", posts)
	})
	http.HandleFunc("/login", login)
	http.HandleFunc("/error", errorP)
	http.HandleFunc("/register", register)

	// ✅ post() est réutilisée ici
	http.HandleFunc("/post", middleware.RequireAuth(post))

	// ✅ category() remplacée par handlers.CategoryHandler
	http.HandleFunc("/category", middleware.RequireAuth(handlers.CategoryHandler))

	http.HandleFunc("/logout", logoutHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
