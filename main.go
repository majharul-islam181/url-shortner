package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {

	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("Hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("Hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Encoding to String : ", hash)
	fmt.Println("hasing to 8 : ", hash[:8])
	return hash[:8]

}

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


func main() {
	fmt.Println("Creating Url-Shortner")
	OriginalURL := "www.facebook.com"

	generateShortURL(OriginalURL)
}
