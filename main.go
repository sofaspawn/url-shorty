// objective 1: start a simple server
package main

import (
	"fmt"
	"net/http"
)

// handles "/" requests
func handleRoot(w http.ResponseWriter, r *http.Request) {
	const redirect string = "use the /shorten endpoint to create a shortened URL"
	// fmt.Println(r) // debugging purposes
	fmt.Fprintf(w, redirect)
}

func handleGetShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use POST /shorten to create a shortened URL")
	fmt.Fprintf(w, "format: {\"url\": \"https://example.com\"}")
}

func handlePostShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST not implemented yet")
}

func main() {
	http.HandleFunc("/", handleRoot)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /shorten", handleGetShorten)
	mux.HandleFunc("POST /shorten", handlePostShorten)

	PORT := ":8080"

	fmt.Println("Server is running on http://localhost" + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}
