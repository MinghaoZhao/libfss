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
  CONN_PORT = 8000
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
  fmt.Println(ans0,ans1)
}

func queryServer(queryType, fssKey, prfKeys, numBits string, serverNum int) int {
  fmt.Println("numBits: ", numBits)
  address := "http://"+CONN_HOST+":"+strconv.Itoa(8000+serverNum)+"/"+queryType+"/"+fssKey+"/"+prfKeys+"/"+numBits
  fmt.Println(address)
  resp, _ := http.Get(address)
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  var answer map[string]string
  _ = json.Unmarshal(body, &answer)
  fmt.Println("answer: ", answer)  
  output, _ := strconv.Atoi(answer["ans"])
  fmt.Println("output: ",output) 
  return output
}

func packageKeys(key interface{}) string {
  marshalledKey, _ := json.Marshal(key)
  return base64.StdEncoding.EncodeToString(marshalledKey)
}