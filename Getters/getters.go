package Getters

import (
    "encoding/json"
    "net/http"
    "os"
    "fmt"

)

// Structure pour décoder les données JSON de l'API
type Artist struct {
    ID     int      `json:"id"`
    Name   string   `json:"name"`
    Date   int      `json:"creationDate"`
    Image  string   `json:"image"`
    Membre []string `json:"members"`
    Albums string   `json:"firstAlbum"`
    SpotifyLink string  `json:"-"`
}

type SpotifyData struct {
    ID          int    `json:"id"`
    SpotifyLink string `json:"spotifyLink"`
}


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
    for i := range artists {
        artists[i].GetSpotifyLink()
    }
    return artists, nil
}

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
    artist.GetSpotifyLink()
    return artist, nil
}

// Fonction pour lire les données à partir du fichier spotify.json
func getSpotifyData() ([]SpotifyData, error) {
    filepath , _ := os.Getwd()
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
