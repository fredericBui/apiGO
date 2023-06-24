package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// Structure pour représenter un article
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

// Liste d'articles
type Articles []Article

// Handler pour l'endpoint POST /articles
func postArticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Post endpoint worked")
}

// Handler pour l'endpoint GET /articles
func allArticles(w http.ResponseWriter, r *http.Request){
	
	// Liste d'exemple d'articles
	articles := Articles{
		Article{Title: "My title", Desc: "My description", Content: "My content"},
		Article{Title: "My title 2", Desc: "My description 2", Content: "My content 2"},
	}

	fmt.Println("Endpoint Hit: all articles")

	// Encodage de la liste d'articles en JSON et envoi dans la réponse
	json.NewEncoder(w).Encode(articles)
}

// Handler pour l'endpoint racine "/"
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

// Configuration des routes et démarrage du serveur
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Association des handlers aux endpoints
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
    
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    handleRequests()
}