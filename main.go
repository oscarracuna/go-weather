package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// TODO: loop to call API and update the weather
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
	go tickers()
	time.Sleep(time.Second * 86400)
}

func tickers() {
	for range time.Tick(time.Second * 120) {
		resp := getResponse()
		fmt.Printf("%.1f°C\n", resp)
		currentTime := time.Now()
		formattedTime := currentTime.Format("Jan 02, 2006 3:04 PM")
		fmt.Println("Current time:", formattedTime)
	}
}

// TODO: find a way to handle errors here
// func that returns an erro and then stops the time.Sleep()
func getResponse() float64 {
	godotenv.Load()
	api_key := os.Getenv("API_KEY")
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
