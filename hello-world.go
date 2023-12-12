package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ROUTE HELLO
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// SERVE the index.html page over
	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	r.Run(":4444")
}
