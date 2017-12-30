package main

import "github.com/gin-gonic/gin"

type Greeting struct {
	Name string
}

func main() {
	r := gin.Default()
	//GET
	r.GET("/home", func(c *gin.Context) {
		c.JSON(
			200,
			Greeting{Name: "HOME"},
		)
	})

	//GET with params
	r.GET("/sub/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(
			200,
			Greeting{Name: name},
		)
	})

	//GET with string query
	r.GET("/query", func(c *gin.Context) {
		value := c.DefaultQuery("name", "HOME")
		c.JSON(
			200,
			Greeting{Name: value},
		)
	})

	//POST string
	r.POST("/form", func(c *gin.Context) {
		value := c.DefaultPostForm("name", "HOME")
		c.JSON(
			200, 
			Greeting{Name: value},
		)
	})

	//POST upload single file
	r.Run()
}