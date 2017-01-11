package main

import (
  "os"
  "gopkg.in/gin-gonic/gin.v1"
  "../libfss"
)

func main() {
    port := os.Getenv("PORT")
    if (port == "") {
      port = "8000"
    }
    r := gin.Default()
    r.GET("/nodes/:key", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "key": c.Param("key"),
            "qtype": "nodes",
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

func evalKey(enc string) {
  var key libfss.FssKeyEq2P
  _ = json.Unmarshal(decodeKey(enc), &key)
  fmt.Println(key)

  // Simulate server: setup needs prfkeys, numbits
  // fServer := libfss.ServerInitialize(fClient.PrfKeys, fClient.NumBits)
  // ans = fServer.EvaluatePF(0, fssKeys[0], 10)
}

func decodeKey(str string) []byte {
  dec, _ := base64.StdEncoding.DecodeString(str)
  return dec
}