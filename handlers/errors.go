package handlers

import "net/http"

type ErrorData struct {
    Code    int    
    Title   string 
    Message string 
}

func RenderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	RenderError(w, http.StatusNotFound, "La page que vous cherchez n'existe pas.")
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	RenderError(w, http.StatusForbidden, "Vous n'avez pas la permission d'accéder à cette page.")
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	RenderError(w, http.StatusInternalServerError, "Une erreur inattendue s'est produite.")
}