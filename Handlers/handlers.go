package Handlers

import (
    "GROUPIE-TRACKER/Getters"
    "html/template"
    "net/http"
    "strings"
    "strconv"
    "sort"
)

// Préparation des templates HTML
var templates = template.Must(template.ParseGlob("templates/*.html"))

// Handler pour la page d'accueil
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "accueil.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Handler pour la page de liste des artistes
func ListeHandler(w http.ResponseWriter, r *http.Request) {
    // Récupération des filtres de la requête
    searchTerm := r.URL.Query().Get("search")
    orderFilter := r.URL.Query().Get("order-filter")
    dateFilter := r.URL.Query().Get("date-filter")
    if dateFilter == "" {
        dateFilter = "default"
    }
    membreFilter := r.URL.Query().Get("membre-filter")

    // Récupération des artistes
    artists, err := Getters.GetArtists()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Filtrage des artistes par terme de recherche
    if searchTerm != "" {
        searchTerm = strings.ToLower(searchTerm)
        var filteredArtists []Getters.Artist
        for _, artist := range artists {
            if strings.Contains(strings.ToLower(artist.Name), searchTerm) {
                filteredArtists = append(filteredArtists, artist)
            } else {
                for _, member := range artist.Membre {
                    memberNameParts := strings.Fields(strings.ToLower(member))
                    for _, namePart := range memberNameParts {
                        if strings.Contains(namePart, searchTerm) {
                            filteredArtists = append(filteredArtists, artist)
                            break
                        }
                    }
                }
            }
        }
        artists = filteredArtists
    }

    // Tri des artistes par filtre d'ordre
    switch orderFilter {
    case "Alphabétique":
        sort.Slice(artists, func(i, j int) bool {
            return artists[i].Name < artists[j].Name
        })
    case "Date Croissante":
        sort.Slice(artists, func(i, j int) bool {
            return artists[i].Date < artists[j].Date
        })
    case "Date Décroissante":
        sort.Slice(artists, func(i, j int) bool {
            return artists[i].Date > artists[j].Date
        })
    }

    // Filtrage des artistes par date
    if dateFilter != "default" {
        var filteredArtists []Getters.Artist
        for _, artist := range artists {
            if artist.Date == 0 {
                continue
            }
            artistDate := strings.Split(strconv.Itoa(artist.Date), "-")
            artistYear, err := strconv.Atoi(artistDate[0])
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // Filtrage par décennie
            switch dateFilter {
            case "2010-2020":
                if artistYear >= 2010 && artistYear <= 2020 {
                    filteredArtists = append(filteredArtists, artist)
                }
            case "2000-2010":
                if artistYear >= 2000 && artistYear <= 2010 {
                    filteredArtists = append(filteredArtists, artist)
                }
            case "1990-2000":
                if artistYear >= 1990 && artistYear <= 2000 {
                    filteredArtists = append(filteredArtists, artist)
                }
            case "1980-1990":
                if artistYear >= 1980 && artistYear <= 1990 {
                    filteredArtists = append(filteredArtists, artist)
                }
            case "1970-1980":
                if artistYear >= 1970 && artistYear <= 1980 {
                    filteredArtists = append(filteredArtists, artist)
                }
            case "1960-1970":
                if artistYear >= 1960 && artistYear <= 1970 {
                    filteredArtists = append(filteredArtists, artist)
                }
            case "1950-1960":
                if artistYear >= 1950 && artistYear <= 1960 {
                    filteredArtists = append(filteredArtists, artist)
                }
            }
        }
        artists = filteredArtists
    }

    // Filtrage des artistes par nombre de membres
    if membreFilter != "" && membreFilter != "default" {
        var filteredArtists []Getters.Artist
        desiredMembreCount, err := strconv.Atoi(membreFilter)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        for _, artist := range artists {
            if len(artist.Membre) == desiredMembreCount {
                filteredArtists = append(filteredArtists, artist)
            }
        }
        artists = filteredArtists
    }

    // Affichage de la page de liste avec les artistes filtrés
    err = templates.ExecuteTemplate(w, "liste.html", artists)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Handler pour la page d'information sur un artiste
func InfoHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    // Récupération de l'artiste par son ID
    artist, err := Getters.GetArtistByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Affichage de la page d'information avec les détails de l'artiste
    err = templates.ExecuteTemplate(w, "info.html", artist)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}