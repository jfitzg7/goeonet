package goeonet

import (
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
  "time"
)


var client = http.Client{Timeout: 5 * time.Second}

func main() {
  GetRecentOpenEvents()
}

func GetRecentOpenEvents() {
  req, err := http.NewRequest("GET", "https://eonet.sci.gsfc.nasa.gov/api/v3/events", nil)

  if err != nil {
    log.Fatal("NewRequest: ", err)
  }

  req.Header.Add("status", "open")
  req.Header.Add("limit", "20")

  resp, err := client.Do(req)

  if err != nil {
    log.Fatal("Do: ", err)
  }

  responseData, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    log.Fatal("ReadAll: ", err)
  }

  fmt.Println(responseData)
}
