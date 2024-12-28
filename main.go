package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/joho/godotenv" gotta hardcode the API key otherwise my programlol
  // won't run when I move it to /bin/ lol
)

// TODO: add icon for weather condition --> condition:icon
type ApiResponse struct {
	Current struct {
		Temperature float64 `json:"temp_c"`
		Condition   struct {
			Text string `json:"text"`
		} `json:"conditon"`
	} `json:"current"`
}

func main() {
	initialResp := getResponse()
	fmt.Printf("%.1f°C\n", initialResp)
}

func getResponse() float64 {
  api_key := "API Key goes here"
  city := "El%20Paso"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", api_key, city)

	resp, err := http.Get(url)
	if err != nil {
		log.Print("Something went wrong fetching data", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Response not OK lol")
	}

	var weather ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Println("Error decoding data")
	}

	currentTemp := weather.Current.Temperature
	return currentTemp
}
