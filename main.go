package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Name    string
	Author  Author
	Generes []Genere
}
type Author struct {
	Name string
	Sex  string
}
type Genere struct {
	Name string
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/books", loggingMiddleware(bookHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, porcelain!")
}
func bookHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{Name: "Star War",
		Author:  Author{Name: "Porcelain", Sex: "Girl"},
		Generes: []Genere{{Name: "Tech"}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %v, path: %v", r.Method, r.URL.Path)
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		next(w, r)
	}
}
