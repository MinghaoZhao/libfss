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
  client := libfss.ClientInitialize(6)
  // Test with if x = 10, evaluate to 2: figure out numbers based on node in the future
  fssKeys := client.GenerateTreePF(10, 2) 

  fmt.Println(fssKeys)
  fmt.Println(client.PrfKeys)
  encodedFssKey0 := packageKeys(fssKeys[0])
  encodedFssKey1 := packageKeys(fssKeys[1])
  encodedPrfKeys := packageKeys(client.PrfKeys)


  /*
  marshalledFssKey, _ := json.Marshal(fssKeys[0])
  encodedFssKey := base64.StdEncoding.EncodeToString([]byte(marshalledFssKey))
  */

  // simulated, on server side:
  var newKey0 libfss.FssKeyEq2P
  _ = json.Unmarshal(decodeKey(encodedFssKey0), &newKey0)
  var newKey1 libfss.FssKeyEq2P
  _ = json.Unmarshal(decodeKey(encodedFssKey1), &newKey1)
  fmt.Println(newKey0, newKey1)
  var newPrfKey [][]byte
  _ = json.Unmarshal(decodeKey(encodedPrfKeys), &newPrfKey)
  fmt.Println(newPrfKey)

/*
  fServer := libfss.ServerInitialize(client.PrfKeys, client.NumBits)
  ans := fServer.EvaluatePF(0, newKey, 10)
  fmt.Println(ans)

  resp, _ := http.Get("http://localhost:8000/"+queryType+"/"+encodedFssKey)

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  var newNode map[string]interface{}
  _ = json.Unmarshal(body, &newNode)
  fmt.Println(node)
*/
  // do again for server 2, port 8001
}

func packageKeys(key interface{}) string {
  marshalledKey, _ := json.Marshal(key)
  return base64.StdEncoding.EncodeToString(marshalledKey)
}



func unpackagePrfKey(str string) {
  dec, _ := base64.StdEncoding.DecodeString(str)
  var newkey libfss.FssKeyEq2P
  _ = json.Unmarshal(dec, &newkey)
}