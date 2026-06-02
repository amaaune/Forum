package handlers

import (
	"forum/database"
	"forum/models"
)

func GetPosts() ([]models.Post, error) {
	rows, err := database.DB.Query("SELECT * FROM posts")
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

func CreatePost(user int, title string, content string) error {
	_, err := database.DB.Exec("INSERT INTO posts (user, title, content) VALUES (?, ?, ?)", user, title, content)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(postID int) error {
	_, err := database.DB.Exec("DELETE FROM posts WHERE post_id = ?", postID)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(postID int, title string, content string) error {
	_, err := database.DB.Exec("UPDATE posts SET title = ?, content = ? WHERE post_id = ?", title, content, postID)
	if err != nil {
		return err
	}
	return nil
}
