package database

import (
	"database/sql"
	"forum/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Impossible de Joindre la DB")
	}
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := DB.QueryRow("SELECT user_id, email, username, password FROM users WHERE email = ?", email).
		Scan(&user.UserID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateTables() {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		username TEXT NOT NULL,
		password TEXT
    );
	CREATE TABLE IF NOT EXISTS posts (
		post_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user INTEGER,
		title TEXT,
		content TEXT,
		created_at DATETIME,
		FOREIGN KEY (user) REFERENCES users (user_id)
	);
	CREATE TABLE IF NOT EXISTS comments (
		comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
		post INTEGER,
		user INTEGER,
		FOREIGN KEY (post) REFERENCES posts (post_id),
		FOREIGN KEY (user) REFERENCES users (user_id)
	);
	CREATE TABLE IF NOT EXISTS likes (
		type TEXT,
		post INTEGER,
		user INTEGER,
		comment INTEGER,
		FOREIGN KEY (post) REFERENCES posts (post_id),
		FOREIGN KEY (user) REFERENCES users (user_id),
		FOREIGN KEY (comment) REFERENCES comments (comment_id)
	);
	CREATE TABLE IF NOT EXISTS categories (
		categorie_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	CREATE TABLE IF NOT EXISTS post_categories (
		post_id INTEGER,
		categorie_id INTEGER,
		FOREIGN KEY (post_id) REFERENCES posts (post_id),
		FOREIGN KEY (categorie_id) REFERENCES categories (categorie_id)
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
