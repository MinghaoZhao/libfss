package main

import (
  "os"
  "fmt"
  "../libfss"
  "strconv"
  "encoding/json"
  "encoding/base64"
  "gopkg.in/gin-gonic/gin.v1"
)

const (
  CONN_HOST = "localhost"
  CONN_START_PORT = 8000
)

func main() {
  port := os.Getenv("PORT")
  if (port == "") {
    port = strconv.Itoa(CONN_START_PORT)
  }
  r := gin.Default()
  r.GET("/0/:fssKey/:prfKeys/:numBits", func(c *gin.Context) {
    ans := evalQuery0(c.Param("fssKey"), c.Param("prfKeys"), c.Param("numBits"), port)
    c.JSON(200, gin.H{
        "ans": ans,
        "qtype": "0",
    })
  })
  r.GET("/1/:fssKey/:prfKeys/:numBits", func(c *gin.Context) {
    ans := evalQuery1(c.Param("fssKey"), c.Param("prfKeys"), c.Param("numBits"), port)
    c.JSON(200, gin.H{
        "ans": ans,
        "qtype": "1",
    })
  })
  r.Run(CONN_HOST + ":" + port)
}

func evalQuery0(fssKey, prfKeys, numBits, port string) string {
  // Initialize fss server with PRF keys and number of bits
  var parsedPrfKeys [][]byte
  _ = json.Unmarshal(decodeKey(prfKeys), &parsedPrfKeys)
  parsedNumBits, _ := strconv.ParseUint(numBits, 10, 32)
  Server := libfss.ServerInitialize(parsedPrfKeys, uint(parsedNumBits))

  // Get server number given the port
  portNum, _ := strconv.Atoi(port)
  serverNum := byte(portNum - CONN_START_PORT)

  // Get FSS Key
  var parsedFssKey libfss.FssKeyEq2P
  _ = json.Unmarshal(decodeKey(fssKey), &parsedFssKey)

  // Evaluate PF over all values of DB
  ans := readFetchSmallValue(Server, serverNum, parsedFssKey)
  fmt.Println("answer: ",ans)
  return ans
}


func evalQuery1(fssKey, prfKeys, numBits, port string) string {
  // Initialize fss server with PRF keys and number of bits
  var parsedPrfKeys [][]byte
  _ = json.Unmarshal(decodeKey(prfKeys), &parsedPrfKeys)
  parsedNumBits, _ := strconv.ParseUint(numBits, 10, 32)
  Server := libfss.ServerInitialize(parsedPrfKeys, uint(parsedNumBits))

  // Get server number given the port
  portNum, _ := strconv.Atoi(port)
  serverNum := byte(portNum - CONN_START_PORT)

  // Get FSS Key
  var parsedFssKey libfss.FssKeyEq2P
  _ = json.Unmarshal(decodeKey(fssKey), &parsedFssKey)

  // Evaluate PF over all values of DB
  ans := readFetchLargeValue(Server, serverNum, parsedFssKey)
  fmt.Println("answer: ",ans)
  return ans
}

func decodeKey(str string) []byte {
  dec, _ := base64.StdEncoding.DecodeString(str)
  return dec
}