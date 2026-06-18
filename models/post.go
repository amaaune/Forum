package models

import (
	"fmt"
	"time"
)

type Post struct {
	PostID     int
	User       int
	Title      string
	Content    string
	CreatedAt  time.Time
	Score      int
	Categories []Category
	UserVote string
}

func (p Post) TimeAgo() string {
	duration := time.Since(p.CreatedAt)

	if duration.Seconds() < 60 {
		return "À l'instant"
	}
	if duration.Minutes() < 60 {
		return fmt.Sprintf("il y a %d min", int(duration.Minutes()))
	}
	if duration.Hours() < 24 {
		return fmt.Sprintf("il y a %d h", int(duration.Hours()))
	}
	days := int(duration.Hours() / 24)
	if days < 7 {
		return fmt.Sprintf("il y a %d j", days)
	}

	// Si ça fait plus d'une semaine, format propre : JJ/MM/AAAA
	return p.CreatedAt.Format("02/01/2006")
}
