package handlers

import (
    "sort"
    "time"
)

type PostScore struct {
    PostID int
    Title  string
    Score  float64
}

func Triepop() []string {

    postes, err := GetPosts()
    if err != nil || len(postes) == 0 {
        return []string{}
    }

    var scores []PostScore

    for _, post := range postes {
        likes, err := GetInteraction(post.PostID)
        if err != nil {
            continue
        }

        elapsed := time.Since(post.CreatedAt).Hours()
        if elapsed < 1 {
            elapsed = 1  
        }

        score := float64(likes) / elapsed

        scores = append(scores, PostScore{
            PostID: post.PostID,
            Title:  post.Title,
            Score:  score,
        })
    }

    sort.Slice(scores, func(i, j int) bool {
        return scores[i].Score > scores[j].Score
    })

    var listefinal []string
    limit := 5
    if len(scores) < limit {
        limit = len(scores)
    }

    for i := 0; i < limit; i++ {
        listefinal = append(listefinal, scores[i].Title)
    }

    return listefinal
}