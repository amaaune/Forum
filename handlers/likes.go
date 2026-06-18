package handlers

import "forum/database"

func LikesPost(postID int, userID int) error {
	var currentCount int
	err := database.DB.QueryRow("SELECT count FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID).Scan(&currentCount)
	
	_, _ = database.DB.Exec("DELETE FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID)

	if err == nil && currentCount == 1 {
		return nil
	}

	_, err = database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (1, ?, ?, NULL)", postID, userID)
	return err
}

func DislikesPost(postID int, userID int) error {
	var currentCount int
	err := database.DB.QueryRow("SELECT count FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID).Scan(&currentCount)

	_, _ = database.DB.Exec("DELETE FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID)

	if err == nil && currentCount == -1 {
		return nil
	}

	_, err = database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (-1, ?, ?, NULL)", postID, userID)
	return err
}

func LikesComment(commentID int, userID int) error {
	_, _ = database.DB.Exec("DELETE FROM likes WHERE comment = ? AND user = ? AND post IS NULL", commentID, userID)
	_, err := database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (1, NULL, ?, ?)", userID, commentID)
	return err
}

func DislikesComment(commentID int, userID int) error {
	_, _ = database.DB.Exec("DELETE FROM likes WHERE comment = ? AND user = ? AND post IS NULL", commentID, userID)
	_, err := database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (-1, NULL, ?, ?)", userID, commentID)
	return err
}

func GetInteraction(postID int) (int, error) {
	row := database.DB.QueryRow(`SELECT COALESCE(SUM(count), 0) FROM likes WHERE post = ?`, postID)
	var counts int
	err := row.Scan(&counts)
	return counts, err
}

func GetCommentInteraction(commentID int) (int, error) {
	row := database.DB.QueryRow(`SELECT COALESCE(SUM(count), 0) FROM likes WHERE comment = ?`, commentID)
	var counts int
	err := row.Scan(&counts)
	return counts, err
}