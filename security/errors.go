package security

func RenderError(w http.ResponseWriter, code int, message string) {

    title := "Erreur"

    switch code {
    case http.StatusNotFound:          
        title = "Page introuvable"   // error 404

    case http.StatusForbidden:       
        title = "Accès interdit"	 // error 403

    case http.StatusInternalServerError: 
        title = "Erreur serveur"	 // error 500

    case http.StatusUnauthorized:       
        title = "Non autorisé"	 	// error 401
    }

   
    data := ErrorData{
        Code:    code,
        Title:   title,
        Message: message,
    }

    
    t, err := template.ParseFiles("templates/error.html")
    if err != nil {
        
        http.Error(w, "Erreur critique du serveur", http.StatusInternalServerError)
        return
    }

    
    w.WriteHeader(code)
    t.Execute(w, data)
}