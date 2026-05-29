package security

import (
    "net/http"
    "github.com/gofrs/uuid/v5"
)

func generateUUID() (string, error) {
    id, err := uuid.NewV4()
    if err != nil {
        return "", err
    }
    return id.String(), nil
}

func storeUUID(w http.ResponseWriter) (string, error) {
    sessionID, err := generateUUID()
    if err != nil {
        return "", err
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    sessionID,
        HttpOnly: true,
        Secure:   true,
    })

    return sessionID, nil
}

func validateUUID(uuidStr string) bool {
    id, err := uuid.FromString(uuidStr)
    if err != nil {
        return false
    }
    return id.Version() == uuid.V4
}
