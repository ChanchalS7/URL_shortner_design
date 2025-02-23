package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string `json:"id"`
	OriginalURL  string `json:"original_url"`
	ShortURL     string `json:"short_url"`
	CreationDate string `json:"creation_date"`
}

// Memorybased db using :
var urlDB = make(map[string]URL)

// method to gnerate the shorturl value
func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("final string:", hash)
	fmt.Println("final string:", hash[:8])
	return hash[:8]

}

// method to create shorturl using above methods
func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL

	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// method to get url in response
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New{"URL not found"}
	}
	return url, nil
}

// method to get rootpage URL
func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// handler method for shorturl
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
}

// redirect method after request
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	fmt.Println("Hello")
}
