package Getters

import (
	"encoding/json"
	"net/http"
)

// Structure pour décoder les données JSON de l'API
type Artist struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Date  int    `json:"creationDate"`
	Image string `json:"image"`
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
	return artists, nil
}
