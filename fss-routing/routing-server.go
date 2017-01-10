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