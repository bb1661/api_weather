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
	latitude              float64 `json:"latitude"`
	longitude             float64
	generationtime_ms     float64
	utc_offset_seconds    float64
	timezone              string
	timezone_abbreviation string
	elevation             float64
	weather               current_weather `json:"current_weather"`
}

// A Pokemon Struct to map every pokemon to.
type current_weather struct {
	temperature   float64 `json:"temperature"`
	windspeed     float64
	winddirection float64
	weathercode   int
	time          string
}

func main() {
	//response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	response, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println([]byte(responseData))
	var responseObject Response
	json.Unmarshal([]byte(responseData), &responseObject)

	fmt.Println(responseObject.latitude)

}
