package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "strconv"
  "github.com/cathieyun/libfss/libfss"
  "encoding/json"
  "encoding/base64"
  "encoding/binary"
  "time"
  "strings"
  "bytes"
  "crypto/sha256"
)

const (
  CONN_HOST = "localhost"
  CONN_START_PORT = 8000
)

func main() {
  // makeQuery(0, 191397, 20)
  makeQuery(3, 10, 20)
}

func makeQuery(queryType int, lookup uint, size uint) string {
  t0 := time.Now()

  // Initialize client and generate keys based on query
  client := libfss.ClientInitialize(size)
  fssKeys := client.GenerateTreePF(lookup, 1) 

  chan0 := make(chan string)
  chan1 := make(chan string)
  go queryServer(chan0, strconv.Itoa(queryType), packageKeys(fssKeys[0]), packageKeys(client.PrfKeys), strconv.Itoa(int(client.NumBits)), 0)
  go queryServer(chan1, strconv.Itoa(queryType), packageKeys(fssKeys[1]), packageKeys(client.PrfKeys), strconv.Itoa(int(client.NumBits)), 1)
  ans0 := <-chan0
  ans1 := <-chan1

  t1 := time.Now()
  fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
  if queryType == 2 {
    int0, _ := strconv.Atoi(ans0)
    int1, _ := strconv.Atoi(ans1)
    fmt.Println("answer for querytype 2: ", int0 + int1)
    return strconv.Itoa(int0 + int1)
  } else if queryType == 3 {
    received0 := strings.Split(ans0,",")
    received1 := strings.Split(ans1,",")
    parsed := make([]byte, len(received0))
    for i := range received0 {
      num0, _ := strconv.Atoi(received0[i])
      num1, _ := strconv.Atoi(received1[i])
      parsed[i] = byte(num0 + num1)
    }
    fmt.Println("answer for querytype 3: ",string(parsed))
    return string(bytes.Trim(parsed, "\x00"))
  } else if queryType == 4 {
    received0 := strings.Split(ans0,",")
    received1 := strings.Split(ans1,",")
    parsed := make([]byte, len(received0))
    for i := range received0 {
      num0, _ := strconv.Atoi(received0[i])
      num1, _ := strconv.Atoi(received1[i])
      parsed[i] = byte(num0 + num1)
    }
    fmt.Println("answer for querytype 4: ",string(parsed))
    return string(bytes.Trim(parsed, "\x00"))    
  }
  return ""
}

func stringToInt(s string) uint {
  h := sha256.New()
  h.Write([]byte(s))
  num := binary.LittleEndian.Uint32(h.Sum(nil))
  return uint(num)
}

func queryServer(c chan string, queryType, fssKey, prfKeys, numBits string, serverNum int) {
  port := strconv.Itoa(CONN_START_PORT+serverNum)
  resp, _ := http.Get("http://"+CONN_HOST+":"+port+"/"+queryType+"/"+fssKey+"/"+prfKeys+"/"+numBits)
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  var answer map[string]string
  _ = json.Unmarshal(body, &answer)
  c <- answer["ans"]
}

func packageKeys(key interface{}) string {
  marshalledKey, _ := json.Marshal(key)
  return base64.StdEncoding.EncodeToString(marshalledKey)
}