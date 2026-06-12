package tests

import (
	"forum/database"
	"forum/handlers"
	"os"
	"testing"
	"time"
)

func TestGetPosts(t *testing.T) {
	os.Chdir("../")
	database.InitDB()
	defer database.DB.Close()
	database.CreateTables()

	_, errc := database.DB.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", "admin@test.com", "admin", 1234)
	if errc != nil {
		t.Fatal(errc)
	}

	_, err := database.DB.Exec("INSERT INTO posts (user, title, content, created_at) VALUES (?, ? , ?, ?)", 1, "Test premier du nom", "Voici le premier post de ce forum ovus pouvez le liker si vous voules marquer l'histoire", time.Now())
	if err != nil {
		t.Fatal(err)
	}

	_, err2 := database.DB.Exec("INSERT INTO posts (user, title, content, created_at) VALUES (?, ?, ?, ?)", 1, "Les meilleures pizzas de Paris", "Je cherche vos recommendations pour les meilleures pizzerias de Paris, j'ai déjà essayé Da Graziella mais je veux découvrir autre chose !", time.Now())
	if err2 != nil {
		t.Fatal(err2)
	}

	_, err3 := database.DB.Exec("INSERT INTO posts (user, title, content, created_at) VALUES (?, ?, ?, ?)", 1, "Apprendre le Go en 2024", "Quelles sont vos ressources préférées pour apprendre Go ? Livres, vidéos, projets pratiques ?", time.Now())
	if err3 != nil {
		t.Fatal(err3)
	}

	value, errv := handlers.GetPosts()
	if errv != nil {
		t.Fatal(errv)
	}
	if len(value) == 0 {
		t.Fatal("Pas assez de posts")
	}
}
