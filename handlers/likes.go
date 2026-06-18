package handlers

import "forum/database"

func LikesPost(postID int, userID int) error {
	// 1. On va d'abord vérifier si le même vote existe déjà pour l'annuler
	var currentCount int
	err := database.DB.QueryRow("SELECT count FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID).Scan(&currentCount)
	
	// Dans tous les cas, on nettoie l'ancien vote
	_, _ = database.DB.Exec("DELETE FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID)

	// Si le vote existant était déjà un upvote (1), l'utilisateur voulait l'annuler -> on s'arrête là !
	if err == nil && currentCount == 1 {
		return nil
	}

	// 2. Sinon, on insère le nouveau upvote
	_, err = database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (1, ?, ?, NULL)", postID, userID)
	return err
}

func DislikesPost(postID int, userID int) error {
	// 1. On vérifie si le même vote existe déjà pour l'annuler
	var currentCount int
	err := database.DB.QueryRow("SELECT count FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID).Scan(&currentCount)

	// Dans tous les cas, on nettoie l'ancien vote
	_, _ = database.DB.Exec("DELETE FROM likes WHERE post = ? AND user = ? AND comment IS NULL", postID, userID)

	// Si le vote existant était déjà un downvote (-1), l'utilisateur voulait l'annuler -> on s'arrête là !
	if err == nil && currentCount == -1 {
		return nil
	}

	// 2. Sinon, on insère le nouveau downvote
	_, err = database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (-1, ?, ?, NULL)", postID, userID)
	return err
}

func LikesComment(commentID int, userID int) error {

	var currentCount int

	err := database.DB.QueryRow(
		`SELECT count
		 FROM likes
		 WHERE comment = ?
		 AND user = ?
		 AND post IS NULL`,
		commentID,
		userID,
	).Scan(&currentCount)

	_, _ = database.DB.Exec(
		`DELETE FROM likes
		 WHERE comment = ?
		 AND user = ?
		 AND post IS NULL`,
		commentID,
		userID,
	)

	if err == nil && currentCount == 1 {
		return nil
	}

	_, err = database.DB.Exec(
		`INSERT INTO likes
		(count, post, user, comment)
		VALUES (1, NULL, ?, ?)`,
		userID,
		commentID,
	)

	return err
}

func DislikesComment(commentID int, userID int) error {

	var currentCount int

	err := database.DB.QueryRow(
		`SELECT count
		 FROM likes
		 WHERE comment = ?
		 AND user = ?
		 AND post IS NULL`,
		commentID,
		userID,
	).Scan(&currentCount)

	_, _ = database.DB.Exec(
		`DELETE FROM likes
		 WHERE comment = ?
		 AND user = ?
		 AND post IS NULL`,
		commentID,
		userID,
	)

	if err == nil && currentCount == -1 {
		return nil
	}

	_, err = database.DB.Exec(
		`INSERT INTO likes
		(count, post, user, comment)
		VALUES (-1, NULL, ?, ?)`,
		userID,
		commentID,
	)

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