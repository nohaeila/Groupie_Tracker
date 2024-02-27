package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
)

// Structure pour décoder les données JSON de l'API
type Artist struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Date    int `json:"creationDate"`
    Image     string `json:"image"`
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/liste", listeHandler) // Gestionnaire pour la page de liste
    fmt.Println("Serveur démarré sur le port 8080")
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "accueil.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func listeHandler(w http.ResponseWriter, r *http.Request) {
    artists, err := getArtists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = templates.ExecuteTemplate(w, "liste.html", artists)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getArtists() ([]Artist, error) {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var artists []Artist
    err = json.NewDecoder(resp.Body).Decode(&artists)
    if err != nil {
        return nil, err
    }
    return artists, nil
}
