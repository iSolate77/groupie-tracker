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
	RelationURL  = baseURL + "/relations"
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

type locationResponse struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
	Dates     []string `json:"datesList"`
}

type location struct {
	Index []locationResponse `json:"index"`
}

type date struct {
	ID    int      `json:"id"`
	Dates []string `json:"Dates"`
}

type relation struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocation"`
}

type RelationsWrapper struct {
	Index []relation `json:"index"`
}

// fetchDataFromAPI retrieves data from the API endpoints
func fetchDataFromAPI(url string) ([]byte, error) {
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

// parseArtistData parses JSON data for artists
func parseArtistData(data []byte) ([]Artist, error) {
	var artists []Artist
	err := json.Unmarshal(data, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func parseLocationsData(data []byte) ([]locationResponse, error) {
	var response location
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Index, nil
}

func parseDatesData(data []byte) (date, error) {
	var dates date
	err := json.Unmarshal(data, &dates)
	if err != nil {
		return date{}, err
	}
	return dates, nil
}

func parseRelationsData(data []byte) ([]relation, error) {
	var relations RelationsWrapper
	err := json.Unmarshal(data, &relations)
	if err != nil {
		return nil, err
	}
	return relations.Index, nil
}

func FetchPaginatedArtists(pageNumber int) ([]Artist, int, error) {
	// Fetch all artists data (modify this to support pagination in the future)
	artistData, err := fetchDataFromAPI(ArtistsURL)
	if err != nil {
		return nil, 0, err
	}

	artists, err := parseArtistData(artistData)
	if err != nil {
		return nil, 0, err
	}

	// For simplicity, let's assume 10 artists per page
	perPage := 15
	totalPages := (len(artists) + perPage - 1) / perPage

	// Calculate start and end indices for slicing
	start := (pageNumber - 1) * perPage
	end := start + perPage
	if end > len(artists) {
		end = len(artists)
	}

	return artists[start:end], totalPages, nil
}

func FetchArtistByID(id string) (Artist, error) {
	artistData, err := fetchDataFromAPI(ArtistsURL + "/" + id)
	if err != nil {
		return Artist{}, err
	}

	artists, err := parseArtistData(artistData)
	if err != nil {
		return Artist{}, err
	}

	return artists[0], nil
}
