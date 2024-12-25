package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ApiResponse struct {
  Current struct {
    Tempe float64 `json:"temp_c"`
    Condition struct {
      Text string `json:"text"`
    } `json:"conditon"`
  } `json:"current"`
}

func main() {
	godotenv.Load()
  api_key := os.Getenv("API_KEY")
  city := "Ciudad%20Juarez"
  url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", api_key, city)
  resp, err := http.Get(url)
  if err != nil {
    log.Print("Something went wrong fetching data")
    return
  }

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    log.Fatalln("Response not OK lol")
    return
  }

  var weather ApiResponse
  if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
    fmt.Println("Error decoding data")
    return
  }

  fmt.Printf("%.1fC\n", weather.Current.Tempe)
}
