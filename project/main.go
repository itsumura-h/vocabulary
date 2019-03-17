package main

import (
	"./controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.String(200, "Hello world")
}

func presenetView(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://dumblepy.site", "http://localhost:8080"},
		//AllowOrigins: []string{"*"},
	}))

	r.LoadHTMLGlob("vue/dist/*.html")

	r.POST("/products/vocabulary/api/searchByWords", Controllers.SearchByWords)
	r.GET("/products/vocabulary", presenetView)

	//r.Run(":8000")
	r.RunUnix("/var/run/vocabulary.sock")
}
