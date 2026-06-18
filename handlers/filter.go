package handlers

import (
	"forum/database"
	"forum/models"
	"strings"
)

func AddCategoryToPost(postID int, categoryID []int) error {
	var err error
	for _, id := range categoryID {
		_, err = database.DB.Exec("INSERT INTO post_categories (post_id, categorie_id) VALUES (?, ?)", postID, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateCategory(name string) (int, error) {
	var value string
	value = strings.ToLower(strings.TrimSpace(name))
	res, err := database.DB.Exec("INSERT INTO categories (name) VALUES (?)", value)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastID), err
}

func GetCategoryIDByName(name string) (int, error) {
	var value string
	value = strings.ToLower(strings.TrimSpace(name))
	row := database.DB.QueryRow("SELECT categorie_id FROM categories WHERE name = ?", value)
	var category int
	err := row.Scan(&category)
	if err != nil {
		return 0, err
	}
	return category, nil
}

func GetPostsByCategoryName(categoryName string) ([]models.Post, error) {
	id, err := GetCategoryIDByName(categoryName)
	if err != nil {
		return nil, err
	}
	res, err := GetPostsByCategory(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetPostCategories(postID int) ([]models.Category, error) {
	rows, err := database.DB.Query(`SELECT c.categorie_id, c.name 
		FROM categories c
		JOIN post_categories pc ON c.categorie_id = pc.categorie_id 
		WHERE pc.post_id = ? ORDER BY name ASC`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.CategoryID, &cat.Name); err != nil {
			return categories, err
		}
		categories = append(categories, cat)
	}
	if err = rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}

func GetAllCategories() ([]models.Category, error) {
	rows, err := database.DB.Query("SELECT * FROM categories ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.CategoryID, &cat.Name); err != nil {
			return categories, err
		}
		categories = append(categories, cat)
	}
	if err = rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}

func GetPostsByUserID(userID int) ([]models.Post, error) {
	rows, err := database.DB.Query(`
        SELECT post_id, user, title, content, created_at 
        FROM posts 
        WHERE user = ?`, userID)
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

func GetLikedPostsByUserID(userID int) ([]models.Post, error) {
	rows, err := database.DB.Query(`
		SELECT DISTINCT
			p.post_id,
			p.user,
			p.title,
			p.content,
			p.created_at
		FROM posts p
		INNER JOIN likes l
			ON p.post_id = l.post
		WHERE l.user = ?
		AND l.count = 1
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err := rows.Scan(
			&post.PostID,
			&post.User,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}

		score, _ := GetInteraction(post.PostID)
		post.Score = score

		cats, _ := GetPostCategories(post.PostID)
		post.Categories = cats

		posts = append(posts, post)
	}

	return posts, rows.Err()
}