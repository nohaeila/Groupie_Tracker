package Handlers

import (
    "GROUPIE-TRACKER/Getters"
    "html/template"
    "net/http"
    "strings"
   
   
)


var templates = template.Must(template.ParseGlob("templates/*.html"))


func IndexHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "accueil.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func ListeHandler(w http.ResponseWriter, r *http.Request) {
    // Récupérer le terme de recherche de la requête
    searchTerm := r.URL.Query().Get("search-bar")

    // Récupérer tous les artistes
    artists, err := Getters.GetArtists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Filtrer les artistes en fonction du terme de recherche
    if searchTerm != "" {
        searchTerm = strings.ToLower(searchTerm)
        var filteredArtists []Getters.Artist
        for _, artist := range artists {
            if strings.Contains(strings.ToLower(artist.Name), searchTerm) {
                filteredArtists = append(filteredArtists, artist)
            }
        }
        artists = filteredArtists
    }

    // Afficher la page avec les artistes filtrés
    err = templates.ExecuteTemplate(w, "liste.html", artists)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")


    artist, err := Getters.GetArtistByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = templates.ExecuteTemplate(w, "info.html", artist)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
