
package main

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"
  "os"
)

type Weather struct {
  CurrentTemperature float32 `json:"tempnow"`
  Timestamp string `json:"timestamp"`
  // rest of the JSON is ignored
}

func main() {
  url := fmt.Sprintf("http://weather.willab.fi/weather.json")
  response, err := http.Get(url)
  if err != nil {
    fmt.Println("Failed to open weather data source")
    os.Exit(3)
  }
  var weather Weather
  err = json.NewDecoder(response.Body).Decode(&weather)
  defer response.Body.Close()
  if err != nil {
    fmt.Println("Unable to parse JSON response")
    os.Exit(4)
  }
  t, err := time.Parse("2006-01-02 15:04 MST", weather.Timestamp)
  if err != nil {
    fmt.Println("Unable to parse timestamp from JSON response")
    os.Exit(5)
  }
  var elapsed = time.Since(t)
  fmt.Printf("%s (%.0f seconds ago)\n", weather.Timestamp, elapsed.Seconds())
  fmt.Printf("Oulu Linnanmaa temperature: %+.2fC\n", weather.CurrentTemperature)

}
