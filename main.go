package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Latitude              float64 `json:"latitude"`
	Longitude             float64
	Generationtime_ms     float64
	Utc_offset_seconds    float64
	Timezone              string
	Timezone_abbreviation string
	Elevation             float64
	Weather               current_weather `json:"current_weather"`
}

// A Pokemon Struct to map every pokemon to.
type current_weather struct {
	Temperature   float64 `json:"temperature"`
	Windspeed     float64
	Winddirection float64
	Weathercode   int
	Time          string
}

func main() {
	//response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	response, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=55.7522&longitude=37.6156&current_weather=true")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
	var responseObject Response
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		panic(err)
	}

	fmt.Println(responseObject.Weather.Temperature)

}
