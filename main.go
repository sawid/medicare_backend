package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! กฟพฟ",
		})
	})
	r.GET("/test_path", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! path",
		})
	})
	r.POST("/test_path_post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! path",
		})
	})
	r.POST("/test_path_post_word", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! path",
		})
	})
	r.POST("/receive_post", func(c *gin.Context) {
	
		

	})
	r.Run()

}