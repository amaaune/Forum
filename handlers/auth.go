package handlers

// import (
// 	"encoding/json"
// 	"forum/database"
// 	"forum/models"
// 	"forum/security"
// 	"net/http"
// )

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type LoginResponse struct {
// 	Success bool        `json:"success"`
// 	Message string      `json:"message"`
// 	SessionID string    `json:"session_id,omitempty"`
// 	UserID  int         `json:"user_id,omitempty"`
// 	Username string     `json:"username,omitempty"`
// }

// func GetUserByEmail(email string) (*models.User, error) {
// 	user := &models.User{}
// 	err := database.DB.QueryRow("SELECT user_id, email, username, password FROM users WHERE email = ?", email).
// 		Scan(&user.UserID, &user.Email, &user.Username, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(LoginResponse{
// 			Success: false,
// 			Message: "Méthode non autorisée",
// 		})
// 		return
// 	}

// 	var req LoginRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(LoginResponse{
// 			Success: false,
// 			Message: "Données invalides",
// 		})
// 		return
// 	}
// 	user, err := GetUserByEmail(req.Email)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(LoginResponse{
// 			Success: false,
// 			Message: "Email ou mot de passe incorrect",
// 		})
// 		return
// 	}
// 	if !security.CheckPassword(req.Password, user.Password) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(LoginResponse{
// 			Success: false,
// 			Message: "Email ou mot de passe incorrect",
// 		})
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	SessionID, err := security.StoreUUID(w)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(LoginResponse{
// 			Success: false,
// 			Message: "Erreur lors de la création de la session",
// 		})
// 		return
// 	}
// 	json.NewEncoder(w).Encode(LoginResponse{
// 		Success:  true,
// 		Message:  "Connexion réussie",
// 		UserID:   user.UserID,
// 		Username: user.Username,
// 		SessionID: SessionID,
// 	})
// }
