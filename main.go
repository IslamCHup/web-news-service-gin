package main

import (
	"github.com/gin-gonic/gin"
	"github.com/IslamCHup/web-news-service-gin/handlers"

)

func main() {

	r := gin.Default()

	r.GET("/", handlers.GetAllNews)
	r.GET("/memory/news/:id", handlers.GetNewsById)
	r.POST("/memory/news", handlers.CreateNews)
	r.PUT("/memory/news/:id", handlers.UpdateNews)
	r.DELETE("/memory/news/:id", handlers.DeleteNewsById)

	r.Run(":8080")

}