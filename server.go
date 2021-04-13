package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello wolrd")
	})
	r.GET("user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "name is %s", name)
	})
	err := r.Run(":8000")
	if err == nil {
		fmt.Println("start failed")
	}
}
