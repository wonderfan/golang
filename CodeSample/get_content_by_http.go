package main

import "fmt"
import "io/ioutil"
import "net/http"

func main(){
    url := "http://www.ask.com"
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
      fmt.Println("The content can not be fetched.")
    }
    resp, err := client.Do(req)
    if err != nil {
      fmt.Println("Error fetching %s: %s", url, err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("Error fetching %s: %s", url, err)
    }
    fmt.Println(fmt.Sprintf("%s", body))
}
