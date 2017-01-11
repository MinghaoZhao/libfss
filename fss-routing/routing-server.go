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

func main() {
    port := os.Getenv("PORT")
    if (port == "") {
      port = "8000"
    }
    r := gin.Default()
    r.GET("/nodes/:fssKey/:prfKeys/:numBits", func(c *gin.Context) {
        ans := evalQuery(c.Param("fssKey"), c.Param("prfKeys"), c.Param("numBits"), port)
        c.JSON(200, gin.H{
            "ans": strconv.Itoa(ans),
            "qtype": "nodes",
            "key": c.Param("fssKey"),
        })
    })
    r.GET("/supernodes/:key", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "key": c.Param("key"),
            "qtype": "supernodes",
        })
    })
    r.Run("localhost:" + port)
}

func evalQuery(FssKey, PrfKeys, NumBits, port string) int {
  // Initialize fss server
  var parsedPrfKeys [][]byte
  _ = json.Unmarshal(decodeKey(PrfKeys), &parsedPrfKeys)
  parsedNumBits, _ := strconv.ParseUint(NumBits, 10, 32)
  Server := libfss.ServerInitialize(parsedPrfKeys, uint(parsedNumBits))
  
  // Evaluate key on server. TODO: evaluate based on port #
  var parsedFssKey libfss.FssKeyEq2P
  _ = json.Unmarshal(decodeKey(FssKey), &parsedFssKey)
  ans := Server.EvaluatePF(0, parsedFssKey, 10)
  fmt.Println("answer: ",ans)
  return ans
}

func decodeKey(str string) []byte {
  dec, _ := base64.StdEncoding.DecodeString(str)
  return dec
}