package handlers

import (
	"forum/database"
	"forum/models"
	"log"
	"net/http"
	"strconv"
)

type IndexRenderData struct {
	Posts      []models.Post
	Categories []models.Category
}

// 1. L'INDEX (La page d'accueil avec injection des scores calculés et statuts de vote)
func IndexHandler(w http.ResponseWriter, r *http.Request, render func(http.ResponseWriter, string, any)) {
	posts, err := GetPosts() // Ta fonction SQL existante
	if err != nil {
		http.Error(w, "Erreur de récupération des posts", http.StatusInternalServerError)
		return
	}

	userID := 1 // ID temporaire de dev en attendant le système de session

	// On boucle pour calculer le vrai score ET le statut du vote de l'utilisateur
	for i := range posts {
		posts[i].UserVote = ""

		realScore, err := GetInteraction(posts[i].PostID)
		if err == nil {
			posts[i].Score = realScore
		}

		var currentVoteValue int
		errVote := database.DB.QueryRow(
			"SELECT count FROM likes WHERE post = ? AND user = ? AND comment IS NULL", 
			posts[i].PostID, 
			userID,
		).Scan(&currentVoteValue)

		if errVote == nil {
			if currentVoteValue == 1 {
				posts[i].UserVote = "up"
			} else if currentVoteValue == -1 {
				posts[i].UserVote = "down"
			}
		}
	}

	categories, err := GetAllCategories()
	if err != nil {
		http.Error(w, "Erreur de récupération des catégories", http.StatusInternalServerError)
		return
	}

	data := IndexRenderData{
		Posts:      posts,
		Categories: categories,
	}

	render(w, "index.html", data)
}

// 2. LE SYSTEME DE VOTE (Intercepte les formulaires des logos A et V)
func VoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	postIDStr := r.FormValue("post_id")
	voteType := r.FormValue("vote_type")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "ID du post invalide", http.StatusBadRequest)
		return
	}

	// ID utilisateur temporaire en attendant ton système de session active
	userID := 1 

	if voteType == "up" {
		err = LikesPost(postID, userID)
	} else if voteType == "down" {
		err = DislikesPost(postID, userID)
	}

	if err != nil {
		log.Println("❌ FORGE LOG - Erreur SQLite :", err, " | PostID :", postID, " | VoteType :", voteType)
		http.Error(w, "Erreur lors du traitement du vote", http.StatusInternalServerError)
		return
	}

	ref := r.Referer()
	if ref == "" {
		ref = "/"
	}
	http.Redirect(w, r, ref, http.StatusSeeOther)
}

func PostHandler(w http.ResponseWriter, r *http.Request, render func(http.ResponseWriter, string, any)) {
	render(w, "post.html", nil)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, render func(http.ResponseWriter, string, any)) {
	render(w, "error.html", nil)
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    // 1. Récupération des champs du formulaire
    postIDStr := r.FormValue("post_id")
    content := r.FormValue("content")

    postID, err := strconv.Atoi(postIDStr)
    if err != nil || content == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // 2. User temporaire (En attendant les sessions, on utilise l'user 1 de test)
    userID := 1

    // 3. Insertion en DB avec ta fonction mise à jour
    err = CreateComment(postID, userID, content)
    if err != nil {
        http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
        return
    }

    // 4. Redirection directe sur le post pour voir son commentaire apparaître !
    http.Redirect(w, r, "/post?id="+postIDStr, http.StatusSeeOther)
}

type PostDetailData struct {
	Post     models.Post
	Comments []models.Comment
}

func SinglePostHandler(w http.ResponseWriter, r *http.Request, render func(http.ResponseWriter, string, any)) {
	// 1. On récupère l'ID présent dans l'URL (ex: /post?id=3)
	idStr := r.URL.Query().Get("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// 2. On appelle ta fonction GetPost pour charger le post
	post, err := GetPost(postID)
	if err != nil {
		// Si le post n'existe pas en DB, on évite la page blanche : direction l'erreur
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	// 3. On charge les catégories associées à ce post pour les badges
	catRows, err := database.DB.Query(`
		SELECT c.categorie_id, c.name 
		FROM categories c
		INNER JOIN post_categories pc ON c.categorie_id = pc.categorie_id
		WHERE pc.post_id = ?`, post.PostID)
	
	if err == nil {
		var postCategories []models.Category
		for catRows.Next() {
			var cat models.Category
			if err := catRows.Scan(&cat.CategoryID, &cat.Name); err == nil {
				postCategories = append(postCategories, cat)
			}
		}
		catRows.Close()
		post.Categories = postCategories
	}

	// 4. On charge les commentaires liés avec ta fonction GetCommentsByPost
	comments, err := GetCommentsByPost(postID)
	if err != nil {
		comments = []models.Comment{} // Tableau vide par sécurité si erreur
	}

	// 5. On assemble le tout et on l'envoie enfin au template !
	data := PostDetailData{
		Post:     post,
		Comments: comments,
	}

	render(w, "post.html", data)
}