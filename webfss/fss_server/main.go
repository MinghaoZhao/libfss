package main

import (
  "os"
  "fmt"
  "github.com/cathieyun/libfss/libfss"
  "strconv"
  "encoding/json"
  "encoding/base64"
  "gopkg.in/gin-gonic/gin.v1"
)

const (
  CONN_HOST = "localhost"
  CONN_START_PORT = 8000
  PRIME1 = 3
  PRIME2 = 5
)

func main() {
  port := os.Getenv("PORT")
  if (port == "") {
    port = strconv.Itoa(CONN_START_PORT)
  }
  r := gin.Default()
  r.GET("/:queryType/:fssKey/:prfKeys/:numBits", func(c *gin.Context) {
    ans := evalQuery(c.Param("queryType"), c.Param("fssKey"), c.Param("prfKeys"), c.Param("numBits"), port)
    c.JSON(200, gin.H{
        "ans": ans,
    })
  })
  r.Run(CONN_HOST + ":" + port)
}

func evalQuery(qtype, fssKey, prfKeys, numBits, port string) string {
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

  ans := ""
  // Evaluate PF over all values of DB, DB file depends on qtype
  if qtype == "0" {
    ans = readOneFetchSmall(Server, serverNum, parsedFssKey, "./routing_data/NY_edge_grid.txt")    
  } else if qtype == "1" {
    ans = readOneFetchLarge(Server, serverNum, parsedFssKey, "./routing_data/NY_zones.txt", 50000)
  } else if qtype == "2" {
    ans = readTwoFetchLarge(Server, serverNum, parsedFssKey, "./routing_data/NY_shortest_paths.txt", 2300)
  } else {
    ans = "invalid query type"
  }
  fmt.Println("answer: ",ans)
  return ans
}

func decodeKey(str string) []byte {
  dec, _ := base64.StdEncoding.DecodeString(str)
  return dec
}