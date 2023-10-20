package api

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

// Artist represents the data structure for artists
type Artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    string              `json:"locations"`
	Concerts     map[string][]string `json:"-"`
	Relations    string              `json:"relations"`
}

type LocationResponse struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
	Dates     []string `json:"datesList"`
}

type Location struct {
	Index []LocationResponse `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"Dates"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocation"`
}

type RelationsWrapper struct {
	Index []Relation `json:"index"`
}

// FetchDataFromAPI retrieves data from the API endpoints
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

// ParseArtistData parses JSON data for artists
func ParseArtistData(data []byte) ([]Artist, error) {
	var artists []Artist
	err := json.Unmarshal(data, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func ParseLocationsData(data []byte) ([]LocationResponse, error) {
	var response Location
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Index, nil
}

func ParseDatesData(data []byte) (Date, error) {
	var dates Date
	err := json.Unmarshal(data, &dates)
	if err != nil {
		return Date{}, err
	}
	return dates, nil
}

func ParseRelationsData(data []byte) ([]Relation, error) {
	var relations RelationsWrapper
	err := json.Unmarshal(data, &relations)
	if err != nil {
		return nil, err
	}
	return relations.Index, nil
}

func FetchPaginatedArtists(pageNumber int) ([]Artist, int, error) {
	// Fetch all artists data (modify this to support pagination in the future)
	artistData, err := FetchDataFromAPI(ArtistsURL)
	if err != nil {
		return nil, 0, err
	}

	artists, err := ParseArtistData(artistData)
	if err != nil {
		return nil, 0, err
	}

	// For simplicity, let's assume 10 artists per page
	perPage := 10
	totalPages := (len(artists) + perPage - 1) / perPage

	// Calculate start and end indices for slicing
	start := (pageNumber - 1) * perPage
	end := start + perPage
	if end > len(artists) {
		end = len(artists)
	}

	return artists[start:end], totalPages, nil
}
