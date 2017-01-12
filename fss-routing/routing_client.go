package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "strconv"
  "../libfss"
  "encoding/json"
  "encoding/base64"
  "time"
)

const (
  CONN_HOST = "localhost"
  CONN_START_PORT = 8000
)

func main() {
  makeQuery("nodes", 191397, 20)
}

func makeQuery(queryType string, lookup uint, size uint) {
  t0 := time.Now()
  // Generate fss Keys on client
  client := libfss.ClientInitialize(size)
  fssKeys := client.GenerateTreePF(lookup, 1) 

  chan0 := make(chan int)
  chan1 := make(chan int)
  go queryServer(chan0, queryType, packageKeys(fssKeys[0]), packageKeys(client.PrfKeys), strconv.Itoa(int(client.NumBits)), 0)
  go queryServer(chan1, queryType, packageKeys(fssKeys[1]), packageKeys(client.PrfKeys), strconv.Itoa(int(client.NumBits)), 1)
  ans0 := <-chan0
  ans1 := <-chan1

  t1 := time.Now()
  fmt.Println("combined answer: ", ans0 + ans1)
  fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

func queryServer(c chan int, queryType, fssKey, prfKeys, numBits string, serverNum int) {
  port := strconv.Itoa(CONN_START_PORT+serverNum)
  resp, _ := http.Get("http://"+CONN_HOST+":"+port+"/"+queryType+"/"+fssKey+"/"+prfKeys+"/"+numBits)
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  var answer map[string]string
  _ = json.Unmarshal(body, &answer)
  output, _ := strconv.Atoi(answer["ans"])
  c <- output
}

func packageKeys(key interface{}) string {
  marshalledKey, _ := json.Marshal(key)
  return base64.StdEncoding.EncodeToString(marshalledKey)
}