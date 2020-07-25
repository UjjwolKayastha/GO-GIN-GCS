package main

import (
	"fmt"
	"go-cloud-storage/cloudbucket"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cloud-storage-bucket", cloudbucket.HandleFileUploadToBucket)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
