package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func JsonArtists() ([]Artist, error) {
	temp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer temp.Body.Close()
	byteResult, _ := io.ReadAll(temp.Body)
	var res []Artist
	err2 := json.Unmarshal(byteResult, &res)
	if err2 != nil {
		return nil, err2
	}
	return res, nil
}

func JsonConcerts(id string) (Concert, error) {
	temp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return Concert{}, err
	}
	defer temp.Body.Close()
	var tour Concert
	byteResult, _ := io.ReadAll(temp.Body)
	err2 := json.Unmarshal(byteResult, &tour)
	if err2 != nil {
		return Concert{}, err2
	}
	return tour, nil
}

func JsonLocations() (Place, error) {
	temp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return Place{}, err
	}
	defer temp.Body.Close()
	var location Place
	byteResult, _ := io.ReadAll(temp.Body)
	err2 := json.Unmarshal(byteResult, &location)
	if err2 != nil {
		return Place{}, err
	}
	return location, nil
}
