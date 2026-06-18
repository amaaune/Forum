package middleware

import (
	"forum/database"
	"forum/security"
	"net/http"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
     return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_id")
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        if !security.IsValidUsername(cookie.Value) {
            DeleteSession(w)
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next(w, r)
    }
}

func DeleteSession(w http.ResponseWriter) {
    http.SetCookie(w, &http.Cookie{
        Name:   "session_id",
        Value:  "",     
        Path:   "/",
        MaxAge: -1,      
    })
}

func GetUserID(r *http.Request) int {
    cookie, err := r.Cookie("session_id")
    if err != nil {
        return 0  // pas de cookie → pas connecté
    }

    var userID int
    err = database.DB.QueryRow(
        "SELECT user_id FROM sessions WHERE session_id = ?", cookie.Value,
    ).Scan(&userID)

    if err != nil {
        return 0 
    }
    return userID
}