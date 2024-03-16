package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/render"
	"log"
	"net/http"
	"strconv"
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

var renderer *render.TemplateReader

func SetRenderer(r *render.TemplateReader) {
	renderer = r
}

// IndexHandler handles the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch page number from query parameters
	pageQuery := r.URL.Query().Get("page")
	pageNumber, _ := strconv.Atoi(pageQuery)
	if pageNumber <= 0 {
		pageNumber = 1
	}

	// Fetch paginated data from API (assuming you have a function for this)
	artists, totalPages, err := api.FetchPaginatedArtists(pageNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	// Fetch the ID from the URL
	id := r.URL.Query().Get("id")

	// Fetch the artist from the API
	artist, err := api.FetchArtistByID(id)
	if err != nil {
		log.Printf("Error fetching artist: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare data for the template
	payload := apiData{
		Title:   artist.Name,
		Artists: []api.Artist{artist},
	}

	// Parse and execute the templates
	err = renderer.Render(r.Context(), w, "base", payload)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
