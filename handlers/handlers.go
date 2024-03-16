package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/render"
	"net/http"
	"strconv"
	"strings"
)

type apiData struct {
	Title       string
	Artists     []api.Artist
	CurrentPage int
	TotalPages  int
	HasNext     bool
	HasPrev     bool
	NextPage    int
	PrevPage    int
}

type locationData struct {
	Title          string
	Artist         api.Artist
	Locations      []string            `json:"locations"`
	DatesLocations map[string][]string `json:"datesLocation"`
}

var renderer *render.TemplateReader

func SetRenderer(r *render.TemplateReader) {
	renderer = r
}

// IndexHandler handles the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch page number from query parameters
	pageQuery := r.URL.Query().Get("page")
	pageNumber, err := strconv.Atoi(pageQuery)
	if err != nil || pageNumber <= 0 {
		pageNumber = 1
	}

	// Fetch paginated data from API (assuming you have a function for this)
	artists, totalPages, err := api.FetchPaginatedArtists(pageNumber)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Calculate start and end indices for slicing
	start := (pageNumber - 1) * 15
	end := start + 15
	if end > len(artists) {
		end = len(artists)
	}

	// Prepare data for the template
	payload := apiData{
		Title:       "Home",
		Artists:     artists,
		CurrentPage: pageNumber,
		TotalPages:  totalPages,
		HasNext:     pageNumber < totalPages,
		HasPrev:     pageNumber > 1,
		NextPage:    pageNumber + 1,
		PrevPage:    pageNumber - 1,
	}

	// Parse and execute the templates
	err = renderer.Render(r.Context(), w, "base", payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	artistID := r.PathValue("id")

	artist, err := api.FetchArtistByID(artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := api.FetchLocationsByID(artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relations, err := api.FetchRelationsByID(artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := locationData{
		Title:          "Location",
		Artist:         artist,
		Locations:      locations.Locations,
		DatesLocations: relations.DatesLocation,
	}

	for location, dates := range payload.DatesLocations {
		correctedLocation := strings.Replace(location, "_", " ", -1)

		if correctedLocation != location {
			payload.DatesLocations[correctedLocation] = dates
			delete(payload.DatesLocations, location)
		}
	}

	err = renderer.Render(r.Context(), w, "location", payload)
}
