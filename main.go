package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var long2shortMap = make(map[string]string)
var short2longMap = make(map[string]string)

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type LongifyResponse struct {
	LongURL string `json:"original_url"`
}

// handles "/" requests
func handleRoot(w http.ResponseWriter, r *http.Request) {
	const redirect string = "use the /shorten endpoint to create a shortened URL"
	// fmt.Println(r) // debugging purposes
	fmt.Fprintf(w, redirect)
}

func handleGetShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use POST /shorten to create a shortened URL\n")
	fmt.Fprintf(w, "format: {\"url\": \"https://example.com\"}")
}

func handlePostShorten(w http.ResponseWriter, r *http.Request) {
	// 1. parse the request body [x]
	// 2. extract the URL from the request body [x]
	// 3. generate a shortened URL using long2short
	// 4. serialize into json
	// 5. return the shortened URL

	var body struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL := long2short(body.URL)

	response := ShortenResponse{
		ShortURL: shortURL,
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(data))
}

func handleGetLongify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use POST /longify to get the original URL\n")
	fmt.Fprintf(w, "format: {\"short_url\": \"short_url\"}")
}

func handlePostLongify(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ShortURL string `json:"short_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longURL, ok := short2longMap[body.ShortURL]
	if !ok {
		http.Error(w, "short URL not found", http.StatusNotFound)
		return
	}

	response := LongifyResponse{
		LongURL: longURL,
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(data))
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("url")
	short_url, ok := long2shortMap[url]
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	delete(long2shortMap, url)
	delete(short2longMap, short_url)

	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /shorten", handleGetShorten)
	mux.HandleFunc("POST /shorten", handlePostShorten)
	mux.HandleFunc("GET /longify", handleGetLongify)
	mux.HandleFunc("POST /longify", handlePostLongify)
	mux.HandleFunc("DELETE /delete/{url}", handleDelete)

	PORT := ":8080"

	fmt.Println("Server is running on http://localhost" + PORT)
	if err := http.ListenAndServe(PORT, mux); err != nil {
		panic(err)
	}
}

func long2short(url string) string {
	val, ok := long2shortMap[url]
	if ok {
		return val
	}
	sum := sha256.Sum256([]byte(url))
	shawty := hex.EncodeToString(sum[:])
	if len(shawty) > 8 {
		shawty = shawty[:8]
	}
	long2shortMap[url] = shawty
	short2longMap[shawty] = url
	return shawty
}
