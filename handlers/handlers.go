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
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locationsData, err := api.FetchDataFromAPI(api.LocationsURL)
	// datesData, err := api.FetchDatesFromAPI()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse data
	artists, err := api.ParseArtistData(artistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locations, err := api.ParseLocationsData(locationsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dates []api.Date
	for i, location := range locations {
		datesData, err := api.FetchDataFromAPI(location.DatesURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		datesForLocation, err := api.ParseDatesData(datesData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		locations[i].Dates = datesForLocation.Dates
		dates = append(dates, datesForLocation)
	}

	relationsData, err := api.FetchDataFromAPI(api.RelationURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	relations, err := api.ParseRelationsData(relationsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	relationMap := make(map[int]map[string][]string)
	for _, relation := range relations {
		relationMap[relation.ID] = relation.DatesLocation
	}

	for i, artist := range artists {
		if datesLocations, ok := relationMap[artist.ID]; ok {
			artists[i].Concerts = datesLocations
		}
	}

	// Render HTML using a template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Artists   []api.Artist
		Locations []api.LocationResponse
		Dates     []api.Date
	}{
		Artists:   artists,
		Locations: locations,
		Dates:     dates,
	}

	// Pass data to the template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
