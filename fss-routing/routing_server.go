package main

import (
  "os"
  "fmt"
  "../libfss"
  "strconv"
  "encoding/json"
  "encoding/base64"
  "gopkg.in/gin-gonic/gin.v1"

  "bufio"
  "strings"
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
    r.GET("/nodes/:fssKey/:prfKeys/:numBits", func(c *gin.Context) {
        ans := evalQuery(c.Param("fssKey"), c.Param("prfKeys"), c.Param("numBits"), port)
        c.JSON(200, gin.H{
            "ans": strconv.Itoa(ans),
            "qtype": "nodes",
        })
    })
    r.GET("/supernodes/:key", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "key": c.Param("key"),
            "qtype": "supernodes",
        })
    })
    r.Run(CONN_HOST + ":" + port)
}

func evalQuery(fssKey, prfKeys, numBits, port string) int {
  // Initialize fss server with PRF keys and number of bits:
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
  ans := readEdgeFile(Server, serverNum, parsedFssKey)
  fmt.Println("server number: ",serverNum)
  fmt.Println("answer: ",ans)
  return ans
}

func decodeKey(str string) []byte {
  dec, _ := base64.StdEncoding.DecodeString(str)
  return dec
}

func readEdgeFile(server *libfss.Fss, serverNum byte, fssKey libfss.FssKeyEq2P) int {
  var ans int = 0
  file, _ := os.Open("./NY_edge_grid.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := strings.Split(scanner.Text(), " ")
    node, _ := strconv.Atoi(line[1])
    weight, _ := strconv.Atoi(line[3])
    ans += server.EvaluatePF(serverNum, fssKey, uint(node))*weight
  }
  fmt.Println(ans)
  return ans
}