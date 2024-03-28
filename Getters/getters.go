package Getters

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "os"
)

// Structures pour décoder les données JSON de l'API
type Artist struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Date        int      `json:"creationDate"`
    Image       string   `json:"image"`
    Membre      []string `json:"members"`
    Albums      string   `json:"firstAlbum"`
    DateConcert []string `json:"-"`  
    Location    []string `json:"-"`
    SpotifyLink string   `json:"-"`
}

type LocationData struct {
    ID       int    `json:"id"`
    Location []string `json:"locations"`
}

type DateData struct {
    ID   int `json:"id"`
    DateConcert []string `json:"dates"`
}

type SpotifyData struct {
    ID          int    `json:"id"`
    SpotifyLink string `json:"spotifyLink"`
}

// Méthode pour obtenir le lien Spotify d'un artiste
func (artist *Artist) GetSpotifyLink() {
    spotifyData, err := getSpotifyData()
    if err != nil {
        return
    }
    for _, data := range spotifyData {
        if data.ID == artist.ID {
            artist.SpotifyLink = data.SpotifyLink
        }
    }
}

// Fonction pour obtenir tous les artistes
func GetArtists() ([]Artist, error) {
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
    // Obtenir les liens Spotify, les dates de concert et les lieux pour chaque artiste
    for i := range artists {
        artists[i].GetSpotifyLink()
        artists[i].GetDateConcert()
        artists[i].GetLocation()
    }
    return artists, nil
}

// Fonction pour obtenir un artiste par son ID
func GetArtistByID(id string) (Artist, error) {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
    if err != nil {
        return Artist{}, err
    }
    defer resp.Body.Close()

    var artist Artist
    err = json.NewDecoder(resp.Body).Decode(&artist)
    if err != nil {
        return Artist{}, err
    }
    // Obtenir le lien Spotify, les dates de concert et les lieux pour l'artiste
    artist.GetDateConcert()
    artist.GetLocation()
    artist.GetSpotifyLink()
    return artist, nil
}

// Méthode pour obtenir les dates de concert d'un artiste
func (artist *Artist) GetDateConcert() {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(artist.ID))
    if err != nil {
        fmt.Println("Erreur lors de la récupération des dates de concert:", err)
        return
    }
    defer resp.Body.Close()

    var dateData DateData
    err = json.NewDecoder(resp.Body).Decode(&dateData)
    if err != nil {
        fmt.Println("Erreur lors du décodage des dates de concert:", err)
        return
    }

    artist.DateConcert = dateData.DateConcert
}

// Méthode pour obtenir les lieux d'un artiste
func (artist *Artist) GetLocation() {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(artist.ID))
    if err != nil {
        return
    }
    defer resp.Body.Close()

    var locationData LocationData
    err = json.NewDecoder(resp.Body).Decode(&locationData)
    if err != nil {
        return
    }

    artist.Location = locationData.Location
}

// Fonction pour lire les données à partir du fichier spotify.json
func getSpotifyData() ([]SpotifyData, error) {
    filepath, _ := os.Getwd()
    fmt.Println(filepath)
    data, err := os.ReadFile(filepath + "/Json/spotify.json")
    if err != nil {
        return nil, err
    }

    var spotifyData []SpotifyData
    err = json.Unmarshal(data, &spotifyData)
    if err != nil {
        return nil, err
    }

    return spotifyData, nil
}