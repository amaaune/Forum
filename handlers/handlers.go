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
	IsAuth     bool
	Username   string
}

func IndexHandler(w http.ResponseWriter, r *http.Request, render func(http.ResponseWriter, string, any)) {
	posts, err := GetPosts()
	if err != nil {
		http.Error(w, "Erreur de récupération des posts", http.StatusInternalServerError)
		return
	}

	userID := 1

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

		catRows, errCat := database.DB.Query(`
			SELECT c.categorie_id, c.name 
			FROM categories c
			INNER JOIN post_categories pc ON c.categorie_id = pc.categorie_id
			WHERE pc.post_id = ?`, posts[i].PostID)
		
		if errCat == nil {
			var postCategories []models.Category
			for catRows.Next() {
				var cat models.Category
				if errCatScan := catRows.Scan(&cat.CategoryID, &cat.Name); errCatScan == nil {
					postCategories = append(postCategories, cat)
				}
			}
			catRows.Close()
			posts[i].Categories = postCategories
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
		IsAuth:     true,
		Username:   "AdminTest",  
	}

	render(w, "index.html", data)
}

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

    postIDStr := r.FormValue("post_id")
    content := r.FormValue("content")

    postID, err := strconv.Atoi(postIDStr)
    if err != nil || content == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    userID := 1

    err = CreateComment(postID, userID, content)
    if err != nil {
        http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/post?id="+postIDStr, http.StatusSeeOther)
}

type PostDetailData struct {
	Post     models.Post
	Comments []models.Comment
}

func SinglePostHandler(w http.ResponseWriter, r *http.Request, render func(http.ResponseWriter, string, any)) {
	idStr := r.URL.Query().Get("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	post, err := GetPost(postID)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

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

	comments, err := GetCommentsByPost(postID)
	if err != nil {
		comments = []models.Comment{}
	}

	data := PostDetailData{
		Post:     post,
		Comments: comments,
	}

	render(w, "post.html", data)
}