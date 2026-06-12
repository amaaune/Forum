package models

import "time"

type Post struct {
	PostID     int
	User       int
	Title      string
	Content    string
	CreatedAt  time.Time
	Score      int
	Categories []Category
}
