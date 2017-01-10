package main

import (
  // "net/http"
  // "io/ioutil"
  "fmt"
  "encoding/json"
  "encoding/base64"
  "../libfss"
)

func main() {
  makeQuery(1234, "nodes")
}

func makeQuery(node int, queryType string) {
  // Generate fss Keys on client
  fClient := libfss.ClientInitialize(6)
  // Test with if x = 10, evaluate to 2
  fssKeys := fClient.GenerateTreePF(10, 2) 
  fmt.Println(fssKeys[0])
  key, _ := json.Marshal(fssKeys[0])
  encoded := base64.StdEncoding.EncodeToString([]byte(key))
  fmt.Println(key)
  fmt.Println(encoded)

  decoded, _ := base64.StdEncoding.DecodeString(encoded)
  var newKey libfss.FssKeyEq2P
  _ = json.Unmarshal(decoded, &newKey)
  fmt.Println(decoded)  
  fmt.Println(newKey)

/*
  resp, err := http.Get("http://localhost:8000/"+queryType+"/")
  if err != nil {
    // handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  var node map[string]interface{}
  err = json.Unmarshal(body, &node)
  fmt.Println(node)

  // do again for server 2, port 8001
*/
}