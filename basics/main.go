package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){ 	
	InitDB()
	defer DB.Close()
	
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	
	e.GET("/", func (c *gin.Context)  {
		todos := ReadToDoList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})
	
	
	e.Run(":8080")
}
