package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "encoding/json"
)

func main() {
  resp, err := http.Get("http://localhost:8000/nodes/1234")
  if err != nil {
    // handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  var node map[string]interface{}
  err = json.Unmarshal(body, &node)
  fmt.Println(node)
}