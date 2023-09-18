package handlers

import (
	"groupie-tracker/api"
	"html/template"
	"net/http"
)

// IndexHandler handles the root URL
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch data from API
	artistData, err := api.FetchDataFromAPI(api.ArtistsURL)
	locationsData, err := api.FetchLocationsFromAPI(api.LocationsURL)
	// datesData, err := api.FetchDatesFromAPI()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse data
	artists, err := api.ParseArtistData(artistData)
	locations, err := api.ParseLocationsData(locationsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render HTML using a template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Artists   []api.Artist
		Locations []api.Location
	}{
		Artists:   artists,
		Locations: locations,
	}

	// Pass data to the template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
