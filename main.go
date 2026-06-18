package main

import (
	"forum/database"
	"forum/handlers"
	"forum/middleware"
	"html/template"
	"log"
	"net/http"
)

// renderTemplate reste ici car elle gère le chemin local vers tes fichiers HTML
// Elle utilise "any" pour accepter n'importe quel type de données (liste de posts, un seul post, ou nil)
func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles("templates/"+tmpl, "templates/poly/header.html", "templates/poly/footer.html", "templates/poly/modal.html")
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

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteSession(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	// 1. Initialisation de la Forge (Base de données)
	database.InitDB()
	defer database.DB.Close()
	database.CreateTables()

	_, _ = database.DB.Exec("INSERT OR IGNORE INTO users (user_id, username, password) VALUES (1, 'AdminTest', 'password_dummy')")

	// 2. Fichiers statiques (CSS, Images, SVG)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// 3. Routes Publiques (On passe renderTemplate en paramètre)
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

	// 4. Routes Protégées par Middleware Auth
	http.HandleFunc("/post", middleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		handlers.PostHandler(w, r, renderTemplate)
	}))

	http.HandleFunc("/post/create", handlers.CreatePostHandler)
	
	http.HandleFunc("/category", middleware.RequireAuth(handlers.CategoryHandler))
	http.HandleFunc("/logout", logoutHandler)

	// 5. La route magique pour intercepter le clic sur les logos A et V 🌟
	http.HandleFunc("/post/vote", handlers.VoteHandler)

	// 6. Lancement du serveur
	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}