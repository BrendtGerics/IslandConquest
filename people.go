package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Person struct {
	Res          map[string]int `json:"resources"`
	Ap           int            `json:"action_points"`
	GranLevel    int            `json:"granary_level"`
	TradingLevel int            `json:"trading_level"`
	Farms        int64          `json:"farms"`
	Mines        int            `json:"mines"`
	Mills        int            `json:"mills"`
	Quarries     int            `json:"quarries"`
	Name         string         `json:"name"`
	Species      string         `json:"species"`
	Description  string         `json:"description"`
}

var peeps []Person

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	peopleListBytes, err := json.Marshal(peeps)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(peopleListBytes)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Person
	person := Person{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	person.Name = r.Form.Get("name")
	person.Description = r.Form.Get("description")
	person.Farms, err = strconv.ParseInt(r.Form.Get("farms"), 0, 10)

	// Append our existing list of birds with a new entry
	peeps = append(peeps, person)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
