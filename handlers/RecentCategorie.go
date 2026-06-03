package handlers

import (
    "forum/database"
    "forum/middleware"
    "html/template"
    "net/http"
    "time"
	"fmt"
)

type RecentCategory struct {
    CategorieID int
    Name        string
    VisitedAt   time.Time
}

func SaveRecentCategory(userID int, categorieID int) error {
    _, err := database.DB.Exec(`
        INSERT INTO recent_categories (user_id, categorie_id, visited_at)
        VALUES (?, ?, datetime('now'))
        ON CONFLICT(user_id, categorie_id)
        DO UPDATE SET visited_at = datetime('now')
    `, userID, categorieID)
    return err
}

func GetRecentCategories(userID int) ([]RecentCategory, error) {
    rows, err := database.DB.Query(`
        SELECT c.categorie_id, c.name, rc.visited_at
        FROM recent_categories rc
        JOIN categories c ON rc.categorie_id = c.categorie_id
        WHERE rc.user_id = ?
        ORDER BY rc.visited_at DESC
        LIMIT 5
    `, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var categories []RecentCategory
    for rows.Next() {
        var cat RecentCategory
        rows.Scan(&cat.CategorieID, &cat.Name, &cat.VisitedAt)
        categories = append(categories, cat)
    }
    return categories, nil
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

    userID := middleware.GetUserID(r)

    categorieID := r.URL.Query().Get("id")
    if categorieID != "" && userID != 0 {
        var id int
        fmt.Sscanf(categorieID, "%d", &id)
        SaveRecentCategory(userID, id)
    }

    var recentCats []RecentCategory
    if userID != 0 {
        recentCats, _ = GetRecentCategories(userID)
    }

    t, err := template.ParseFiles("templates/category.html")
    if err != nil {
        RenderError(w, http.StatusInternalServerError, "Erreur lors du chargement de la page.")
        return
    }

    t.Execute(w, map[string]interface{}{
        "IsLoggedIn":       userID != 0,
        "RecentCategories": recentCats,
    })
}