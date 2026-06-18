package security

import (
    "net/http"
    "time"
    "github.com/gofrs/uuid/v5"
)

func GenerateUUID() (string, error) {
    id, err := uuid.NewV4()
    if err != nil {
        return "", err
    }
    return id.String(), nil
}

func StoreUUID(w http.ResponseWriter) (string, error) {
    sessionID, err := GenerateUUID()
    if err != nil {
        return "", err
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    sessionID,
        HttpOnly: true,
        Secure:   true,
        Path:     "/",             
        Expires:  time.Now().Add(7 * 24 * time.Hour), 
        SameSite: http.SameSiteLaxMode,
    })

    return sessionID, nil
}

func ValidateUUID(uuidStr string) bool {
    id, err := uuid.FromString(uuidStr)
    if err != nil {
        return false
    }
    return id.Version() == uuid.V4
}
