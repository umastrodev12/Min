package pluginweb

import (
	"github.com/gin-gonic/gin"
)

func routes() {
	r := gin.Default() // for routes

	// routes
	r.GET("/version", func(c* gin.Context) {
		c.JSON(200, gin.H{
			"version": "0.0.1 BETA",
			"author": "Astro",
		})
	})

	r.POST("/api/data", func(c* gin.Context) {
		var data struct {
			Data   string `json:"data"`
		}

		c.JSON(200, gin.H{"data": data})
	})

	r.Run(":8080") // start server
}
