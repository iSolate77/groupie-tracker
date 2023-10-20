package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/render"
	"net/http"
	"strconv"
)

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
	data := struct {
		Title       string
		Artists     []api.Artist
		CurrentPage int
		TotalPages  int
		HasNext     bool
		HasPrev     bool
		NextPage    int
		PrevPage    int
	}{
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
	err = renderer.Render(r.Context(), w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
