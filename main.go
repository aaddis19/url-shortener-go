package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type URLShortener struct {
	urls map{string}string
}

func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    originalURL := r.FormValue("url")
    if originalURL == "" {
        http.Error(w, "URL parameter is missing", http.StatusBadRequest)
        return
    }

    shortKey := generateShortKey()
    us.urls[shortKey] = originalURL

    shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

    w.Header().Set("Content-Type", "text/html")
    responseHTML := fmt.Sprintf(`
        <h2>URL Shortener</h2>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
        <form method="post" action="/shorten">
            <input type="text" name="url" placeholder="Enter a URL">
            <input type="submit" value="Shorten">
        </form>
    `, originalURL, shortenedURL, shortenedURL)
    fmt.Fprintf(w, responseHTML)
}