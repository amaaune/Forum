package handlers

import (
	"forum/database"
	"forum/models"
	"net/http"
	"strconv"
)

func GetPosts() ([]models.Post, error) {
	rows, err := database.DB.Query("SELECT * FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostID, &post.User, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return posts, err
	}
	return posts, nil
}

func GetPostsByCategory(category int) ([]models.Post, error) {
	rows, err := database.DB.Query(`SELECT posts.* 
		FROM posts
		INNER JOIN post_categories ON posts.post_id = post_categories.post_id
		WHERE post_categories.categorie_id = ?`, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostID, &post.User, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return posts, err
	}
	return posts, nil
}

func GetPost(postID int) (models.Post, error) {
	row := database.DB.QueryRow("SELECT * FROM posts WHERE post_id = ?", postID)
	var post models.Post
	err := row.Scan(&post.PostID, &post.User, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return post, err
	}
	return post, nil
}

func CreatePost(user int, title string, content string, categoryIDs []int) error {
	res, err := database.DB.Exec("INSERT INTO posts (user, title, content, created_at) VALUES (?, ?, ?, datetime('now'))", user, title, content)
	if err != nil {
        return err
    }
	lastID, err := res.LastInsertId()
    if err != nil {
        return err
    }
    postID := int(lastID)

	for _, catID := range categoryIDs {
        _, err := database.DB.Exec("INSERT INTO post_categories (post_id, categorie_id) VALUES (?, ?)", postID, catID)
        if err != nil {
            return err
        }
    }

    return nil
}

func DeletePost(postID int) error {
	_, err := database.DB.Exec("DELETE FROM posts WHERE post_id = ?", postID)
	return err
}

func UpdatePost(postID int, title string, content string) error {
	_, err := database.DB.Exec("UPDATE posts SET title = ?, content = ? WHERE post_id = ?", title, content, postID)
	return err
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    // 1. Récupération des données du formulaire
    title := r.FormValue("title")
    content := r.FormValue("content")
    
    // r.Form["categories"] récupère TOUTES les valeurs des checkboxes cochées qui ont name="categories"
    r.ParseForm()
    categoryStrings := r.Form["categories"]
    var categoryIDs []int

    // Conversion des IDs de string vers int
    for _, catStr := range categoryStrings {
        if id, err := strconv.Atoi(catStr); err == nil {
            categoryIDs = append(categoryIDs, id)
        }
    }

    // 2. Récupération de l'user (En attendant l'étape 3 des sessions, on utilise notre userID 1 en dur)
    userID := 1 

    err := CreatePost(userID, title, content, categoryIDs)
    if err != nil {
        http.Error(w, "Erreur lors de la création du post: "+err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}