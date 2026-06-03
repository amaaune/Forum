package handlers

import "forum/database"

func LikesPost(postID int, userID int) error {
	_, err := database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (? , ?, ?, NULL) ON CONFLICT (post, user, comment) DO UPDATE SET count = 1", 1, postID, userID)
	return err
}

func LikesComment(commentID int, userID int) error {
	_, err := database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (? , NULL, ?, ?) ON CONFLICT (post, user, comment) DO UPDATE SET count = 1", 1, userID, commentID)
	return err
}

func DislikesPost(postID int, userID int) error {
	_, err := database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (? , ?, ?, NULL) ON CONFLICT (post, user, comment) DO UPDATE SET count = -1", -1, postID, userID)
	return err
}

func DislikesComment(commentID int, userID int) error {
	_, err := database.DB.Exec("INSERT INTO likes (count, post, user, comment) VALUES (? , NULL, ?, ?) ON CONFLICT (post, user, comment) DO UPDATE SET count = -1", -1, userID, commentID)
	return err
}

func GetInteraction(postID int) (int, error) {
	row := database.DB.QueryRow(`SELECT COALESCE(SUM(count), 0)
	FROM likes
	WHERE post = ?`, postID)
	var counts int
	err := row.Scan(&counts)
	return counts, err

}

func GetCommentInteraction(commentID int) (int, error) {
	row := database.DB.QueryRow(`SELECT COALESCE(SUM(count), 0)
	FROM likes
	WHERE comment = ?`, commentID)
	var counts int
	err := row.Scan(&counts)
	return counts, err

}
