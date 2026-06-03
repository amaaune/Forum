package handlers

import (
	"encoding/json"
	"forum/database"
	"forum/security"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	UserID  int         `json:"user_id,omitempty"`
	Username string     `json:"username,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Méthode non autorisée",
		})
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Données invalides",
		})
		return
	}
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Email ou mot de passe incorrect",
		})
		return
	}
	if !security.CheckPassword(req.Password, user.Password) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Email ou mot de passe incorrect",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{
		Success:  true,
		Message:  "Connexion réussie",
		UserID:   user.UserID,
		Username: user.Username,
	})
}
