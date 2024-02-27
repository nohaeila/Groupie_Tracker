// handlers.go

package Handlers

import (
    "GROUPIE-TRACKER/Getters"
    "html/template"
    "net/http"
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
    artists, err := Getters.GetArtists()
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


