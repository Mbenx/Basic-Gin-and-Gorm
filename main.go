package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/article/:title", getArticle)
	router.POST("/article", postArticle)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Golang",
	})
}

func getArticle(c *gin.Context) {
	title := c.Param("title")
	c.JSON(200, gin.H{
		"message": "Welcome to Golang",
		"title":   title,
	})
}

func postArticle(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	c.JSON(200, gin.H{
		"message": "Post Berhasil",
		"title":   title,
		"desc":    desc,
	})
}
