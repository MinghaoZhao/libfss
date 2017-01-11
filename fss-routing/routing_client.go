package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "strconv"
  "../libfss"
  "encoding/json"
  "encoding/base64"
)

const (
  CONN_HOST = "localhost"
  CONN_START_PORT = 8000
)

func main() {
  makeQuery(1234, "nodes")
}

func makeQuery(node int, queryType string) {
  // Generate fss Keys on client
  client := libfss.ClientInitialize(6)
  // Test with if x = 10, evaluate to 2: figure out numbers based on node in the future
  fssKeys := client.GenerateTreePF(10, 2) 

  ans0 := queryServer(queryType, packageKeys(fssKeys[0]), packageKeys(client.PrfKeys), strconv.Itoa(int(client.NumBits)), 0)
  ans1 := queryServer(queryType, packageKeys(fssKeys[1]), packageKeys(client.PrfKeys), strconv.Itoa(int(client.NumBits)), 1)
  fmt.Println("combined answer: ", ans0 + ans1)
}

func queryServer(queryType, fssKey, prfKeys, numBits string, serverNum int) int {
  port := strconv.Itoa(CONN_START_PORT+serverNum)
  resp, _ := http.Get("http://"+CONN_HOST+":"+port+"/"+queryType+"/"+fssKey+"/"+prfKeys+"/"+numBits)
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  var answer map[string]string
  _ = json.Unmarshal(body, &answer)
  output, _ := strconv.Atoi(answer["ans"])
  return output
}

func packageKeys(key interface{}) string {
  marshalledKey, _ := json.Marshal(key)
  return base64.StdEncoding.EncodeToString(marshalledKey)
}