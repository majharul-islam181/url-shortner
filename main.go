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

func getURl(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil

}



func main() {
	fmt.Println("Creating Url-Shortner")
	OriginalURL := "www.facebook.com"

	generateShortURL(OriginalURL)

	// Start the HTTP server on port 3000
	fmt.Println("Starting server on port 3000")
	error := http.ListenAndServe(":3000",nil)
	if error!= nil{
		fmt.Println("Error on starting server")
	}
}
