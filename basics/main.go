package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){ 	
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	
	e.GET("/", func (c *gin.Context)  {
		// c.JSON(200, gin.H{
		// 	"name": "Awesome",
		// })
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "Nicolen",
		})
	})
	
	
	e.Run(":8080")
}
