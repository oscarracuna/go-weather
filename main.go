package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
  api_key := os.Getenv("API_KEY")
  city := "Ciudad Juarez"
  url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", api_key, city)
  resp, err := http.Get(url)
  if err != nil {
    log.Print("Something went wrong fetching data")
    return
  } else {
    fmt.Println(resp)
  }
  defer resp.Body.Close()
}
