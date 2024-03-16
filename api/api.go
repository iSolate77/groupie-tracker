package api

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL      = "https://groupietrackers.herokuapp.com/api"
	ArtistsURL   = baseURL + "/artists"
	LocationsURL = baseURL + "/locations"
	DatesURL     = baseURL + "/dates"
	RelationURL  = baseURL + "/relation"
)

// Artist represents the data structure for artists
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Concerts     string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
	// Dates     []string `json:"datesList"`
}

type date struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}

func FetchDataFromAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func parseArtistData(data []byte) (artists []Artist, err error) {
	err = json.Unmarshal(data, &artists)
	if err != nil {
		return artists, err
	}
	return artists, nil
}

func parseLocationsData(data []byte) (locations Location, err error) {
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}

func parseDatesData(data []byte) (dates []date, err error) {
	err = json.Unmarshal(data, &dates)
	if err != nil {
		return dates, err
	}
	return dates, nil
}

func parseRelationsData(data []byte) (relations Relation, err error) {
	err = json.Unmarshal(data, &relations)
	if err != nil {
		return relations, err
	}
	return relations, nil
}

func FetchPaginatedArtists(pageNumber int) (artists []Artist, pages int, err error) {
	// Fetch all artists data (modify this to support pagination in the future)
	artistData, err := FetchDataFromAPI(ArtistsURL)
	if err != nil {
		return artists, 0, err
	}

	artists, err = parseArtistData(artistData)
	if err != nil {
		return artists, 0, err
	}

	// For simplicity, let's assume 10 artists per page
	perPage := 15
	pages = (len(artists) + perPage - 1) / perPage

	start := (pageNumber - 1) * perPage
	end := start + 15
	if end > len(artists) {
		end = len(artists)
	}

	return artists[start:end], pages, nil
}

func FetchArtistByID(id string) (artist Artist, err error) {
	artistData, err := FetchDataFromAPI(ArtistsURL + "/" + id)
	if err != nil {
		return Artist{}, err
	}

	err = json.Unmarshal(artistData, &artist)
	if err != nil {
		return Artist{}, err
	}

	return artist, nil
}

func FetchLocationsByID(id string) (location Location, err error) {
	locationData, err := FetchDataFromAPI(LocationsURL + "/" + id)
	if err != nil {
		return Location{}, err
	}

	location, err = parseLocationsData(locationData)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}

func FetchRelationsByID(id string) (Relation, error) {
	relationData, err := FetchDataFromAPI(RelationURL + "/" + id)
	if err != nil {
		return Relation{}, err
	}

	relations, err := parseRelationsData(relationData)
	if err != nil {
		return Relation{}, err
	}

	return relations, nil
}
