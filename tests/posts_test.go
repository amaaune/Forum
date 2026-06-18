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

	// 1. Injection de l'utilisateur de test
	_, errc := database.DB.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", "admin@test.com", "admin", "1234")
	if errc != nil {
		t.Fatal(errc)
	}

	// 2. Injection des 3 catégories de base
	categories := []string{"Golang", "Nantes", "Gaming"}
	for i, catName := range categories {
		_, errCat := database.DB.Exec("INSERT INTO categories (categorie_id, name) VALUES (?, ?)", i+1, catName)
		if errCat != nil {
			t.Fatal("Erreur insertion catégorie test:", errCat)
		}
	}

	// 3. Injection des Posts (7 posts au total)
	// Post 1
	_, err1 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ? , ?, ?)", 1, 1, "Test premier du nom", "Voici le premier post de ce forum vous pouvez le liker si vous voulez marquer l'histoire", time.Now().Add(-7*24*time.Hour))
	if err1 != nil {
		t.Fatal(err1)
	}

	// Post 2
	_, err2 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ?, ?, ?)", 2, 1, "Les meilleures pizzas de Paris", "Je cherche vos recommandations pour les meilleures pizzerias de Paris, j'ai déjà essayé Da Graziella mais je veux découvrir autre chose !", time.Now().Add(-6*24*time.Hour))
	if err2 != nil {
		t.Fatal(err2)
	}

	// Post 3
	_, err3 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ?, ?, ?)", 3, 1, "Apprendre le Go en 2026", "Quelles sont vos ressources préférées pour apprendre Go ? Livres, vidéos, projets pratiques ?", time.Now().Add(-5*24*time.Hour))
	if err3 != nil {
		t.Fatal(err3)
	}

	// Post 4 (Nouveau)
	_, err4 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ?, ?, ?)", 4, 1, "Balade à vélo sur les bords de l'Erdre", "Superbe météo aujourd'hui à Nantes pour une sortie cycliste. Des motivés pour une boucle de 15km ce week-end ?", time.Now().Add(-4*24*time.Hour))
	if err4 != nil {
		t.Fatal(err4)
	}

	// Post 5 (Nouveau)
	_, err5 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ?, ?, ?)", 5, 1, "Config serveur Nginx & WSL2", "Quelqu'un a déjà galéré à configurer un reverse-proxy Nginx sous Windows avec WSL2 pour son projet de DevSecOps ?", time.Now().Add(-3*24*time.Hour))
	if err5 != nil {
		t.Fatal(err5)
	}

	// Post 6 (Nouveau)
	_, err6 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ?, ?, ?)", 6, 1, "Le Main Support en sueur sur LoL", "Après le dernier patch, vous jouez quoi en Support ou ADC ? J'ai l'impression que la meta a complètement changé.", time.Now().Add(-2*24*time.Hour))
	if err6 != nil {
		t.Fatal(err6)
	}

	// Post 7 (Nouveau)
	_, err7 := database.DB.Exec("INSERT INTO posts (post_id, user, title, content, created_at) VALUES (?, ?, ?, ?, ?)", 7, 1, "Astrophotographie : premier test au Canon 90D", "Des conseils pour le traitement des fichiers RAW sous Darktable pour faire ressortir les nébuleuses ?", time.Now().Add(-1*24*time.Hour))
	if err7 != nil {
		t.Fatal(err7)
	}

	// 4. Liaison des catégories avec variations (Table post_categories)
	// Légende IDs : 1 = Golang, 2 = Nantes, 3 = Gaming
	links := []struct {
		postID int
		catID  int
	}{
		{1, 2}, // Post 1 a seulement la catégorie "Nantes"
		
		{3, 1}, // Post 3 a la catégorie "Golang"
		
		{4, 2}, // Post 4 a "Nantes"
		
		{5, 1}, // Post 5 a "Golang"
		{5, 2}, // Post 5 a AUSSI "Nantes" (Mix de 2)
		
		{6, 3}, // Post 6 a "Gaming"
		
		{7, 1}, // Post 7 cumule les 3 catégories pour le test ultime ! (Mix de 3)
		{7, 2}, 
		{7, 3},
		// Note : Le Post 2 n'a aucune catégorie liée (Mix de 0)
	}

	for _, link := range links {
		_, errLink := database.DB.Exec("INSERT INTO post_categories (post_id, categorie_id) VALUES (?, ?)", link.postID, link.catID)
		if errLink != nil {
			t.Fatal("Erreur liaison post_categories:", errLink)
		}
	}

	// 5. Exécution de la vérification du Handler
	value, errv := handlers.GetPosts()
	if errv != nil {
		t.Fatal(errv)
	}
	if len(value) != 7 {
		t.Fatalf("Attendu 7 posts, mais obtenu %d", len(value))
	}
}
