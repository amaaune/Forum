package models

import "time"

type Comment struct {
	CommentID int
	Post      int
	User      int
	Content   string
	CreatedAt time.Time
}
