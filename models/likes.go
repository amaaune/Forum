package models

import "time"

type Like struct {
	Count     int
	Post      int
	User      int
	Comment   int
	CreatedAt time.Time
}
