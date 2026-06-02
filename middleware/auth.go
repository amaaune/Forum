package middleware

import (
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