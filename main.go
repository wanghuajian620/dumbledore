package main

import "github.com/gin-gonic/gin"

func main() {
 r := gin.Default() // := 是简短声明 相当于 var r = gin.Default()
 r.GET("/ping", func(c *gin.Context) { // *gin 是什么，*gin.Context是参数类型
  c.JSON(200, gin.H{
   "message": "pong",
  })
 })
 r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
