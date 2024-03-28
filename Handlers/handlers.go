package Handlers

import (
	"GROUPIE-TRACKER/Getters"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
)
// Précompilation des templates HTML pour une utilisation ultérieure
var templates = template.Must(template.ParseGlob("templates/*.html"))
// Constantes pour les différents types de tri
const (
	OrderAlphabetic     = "Alphabétique"
	OrderDateAscending  = "Date Croissante"
	OrderDateDescending = "Date Décroissante"
)

// IndexHandler gère la requête à la page d'accueil
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "accueil.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
// ListeHandler gère la requête à la page de liste des artistes
func ListeHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération des paramètres de requête
	searchTerm := r.URL.Query().Get("search")
	orderFilter := r.URL.Query().Get("order-filter")
	dateFilter := r.URL.Query().Get("date-filter")
	membreFilter := r.URL.Query().Get("membre-filter")
	albumFilter := r.URL.Query().Get("album-filter")
	artists, err := Getters.GetArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Filtrage des artistes en fonction des paramètres de requête
	artists = filterArtists(artists, searchTerm, orderFilter, dateFilter, albumFilter, membreFilter)
	// Envoi de la liste filtrée des artistes à la page de liste
	if err := templates.ExecuteTemplate(w, "liste.html", artists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// filterArtists filtre la liste des artistes en fonction des différents paramètres
func filterArtists(artists []Getters.Artist, searchTerm, orderFilter, dateFilter, albumFilter, membreFilter string) []Getters.Artist {
	artists = filterBySearchTerm(artists, searchTerm)
	artists = sortByOrderFilter(artists, orderFilter)
	artists = filterByDate(artists, dateFilter)
	artists = filterByMembre(artists, membreFilter)
	artists = filterByAlbum(artists, albumFilter)
	return artists
}

// filterBySearchTerm filtre la liste des artistes en fonction du terme de recherche
func filterBySearchTerm(artists []Getters.Artist, searchTerm string) []Getters.Artist {
	if searchTerm == "" {
		return artists
	}

	searchTerm = strings.ToLower(searchTerm)
	var filteredArtists []Getters.Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchTerm) {
			filteredArtists = append(filteredArtists, artist)
		} else {
			for _, member := range artist.Membre {
				if strings.Contains(strings.ToLower(member), searchTerm) {
					filteredArtists = append(filteredArtists, artist)
					break
				}
			}
		}
	}
	return filteredArtists
}

// sortByOrderFilter trie la liste des artistes en fonction du filtre de tri
func sortByOrderFilter(artists []Getters.Artist, orderFilter string) []Getters.Artist {
	switch orderFilter {
	case OrderAlphabetic:
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Name < artists[j].Name
		})
	case OrderDateAscending:
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Date < artists[j].Date
		})
	case OrderDateDescending:
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Date > artists[j].Date
		})
	}
	return artists
}

// filterByDate filtre la liste des artistes en fonction du filtre de date
func filterByDate(artists []Getters.Artist, dateFilter string) []Getters.Artist {
	if dateFilter == "" || dateFilter == "default" {
		return artists
	}

	var filteredArtists []Getters.Artist
	for _, artist := range artists {
		if artist.Date == 0 {
			continue
		}

		artistYear, _ := strconv.Atoi(strconv.Itoa(artist.Date)[:4])
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
	return filteredArtists
}

// filterByAlbum filtre la liste des artistes en fonction du filtre d'album
func filterByAlbum(artists []Getters.Artist, albumFilter string) []Getters.Artist {
	if albumFilter == "" || albumFilter == "default" {
		return artists
	}

	var filteredArtists []Getters.Artist
	for _, artist := range artists {
		if len(artist.Albums) == 0 {
			continue
		}

		splitAlbumDate := strings.Split(artist.Albums, "-")
		if len(splitAlbumDate) < 3 {
			continue // Si la date de l'album n'est pas dans le format attendu, ignorez cet artiste
		}
		albumYear, _ := strconv.Atoi(splitAlbumDate[2])
		switch albumFilter {
		case "2010-2020":
			if albumYear >= 2010 && albumYear <= 2020 {
				filteredArtists = append(filteredArtists, artist)
			}
		case "2000-2010":
			if albumYear >= 2000 && albumYear <= 2010 {
				filteredArtists = append(filteredArtists, artist)
			}
		case "1990-2000":
			if albumYear >= 1990 && albumYear <= 2000 {
				filteredArtists = append(filteredArtists, artist)
			}
		case "1980-1990":
			if albumYear >= 1980 && albumYear <= 1990 {
				filteredArtists = append(filteredArtists, artist)
			}
		case "1970-1980":
			if albumYear >= 1970 && albumYear <= 1980 {
				filteredArtists = append(filteredArtists, artist)
			}
		case "1960-1970":
			if albumYear >= 1960 && albumYear <= 1970 {
				filteredArtists = append(filteredArtists, artist)
			}
		case "1950-1960":
			if albumYear >= 1950 && albumYear <= 1960 {
				filteredArtists = append(filteredArtists, artist)
			}
		}
	}
	return filteredArtists
}

// filterByMembre filtre la liste des artistes en fonction du filtre de membre
func filterByMembre(artists []Getters.Artist, membreFilter string) []Getters.Artist {
	if membreFilter == "" || membreFilter == "default" {
		return artists
	}

	var filteredArtists []Getters.Artist
	desiredMembreCount, _ := strconv.Atoi(membreFilter)
	for _, artist := range artists {
		if len(artist.Membre) == desiredMembreCount {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	return filteredArtists
}

// InfoHandler gère la requête à la page d'information sur un artiste
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	artist, err := Getters.GetArtistByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Envoi des informations sur l'artiste à la page d'information
	if err := templates.ExecuteTemplate(w, "info.html", artist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
