package models

import "time"

type Like struct {
	Type      string
	Post      int
	User      int
	Comment   int
	CreatedAt time.Time
}
