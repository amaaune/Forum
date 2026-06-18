package handlers

import (
	"forum/database"
	"forum/models"
)

func GetCommentsByPost(postID int) ([]models.Comment, error) {
	rows, err := database.DB.Query("SELECT comment_id, post, user, content, created_at FROM comments WHERE post = ? ORDER BY created_at ASC", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.CommentID, &comment.Post, &comment.User, &comment.Content, &comment.CreatedAt); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func CreateComment(postID int, userID int, content string) error {
	_, err := database.DB.Exec("INSERT INTO comments (post, user, content, created_at) VALUES (?, ?, ?, datetime('now'))", postID, userID, content)
	return err
}

func DeleteComment(commentID int) error {
	_, err := database.DB.Exec("DELETE FROM comments WHERE comment_id = ?", commentID)
	return err
}

func UpdateComment(commentID int, content string) error {
	_, err := database.DB.Exec("UPDATE comments SET content = ? WHERE comment_id = ?", content, commentID)
	return err
}
