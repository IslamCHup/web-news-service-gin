package main

import (
	"github.com/gin-gonic/gin"
	"github.com/IslamCHup/web-news-service-gin/handlers"

)

func main() {

	r := gin.Default()

	r.GET("/", handlers.GetAllNews)
	r.GET("/memory/NewsAll/:id", handlers.GetNewsById)
	r.POST("/memory/NewsAll", handlers.CreateNews)
	r.PUT("/memory/NewsAll/:id", handlers.UpdateNews)
	r.DELETE("/memory/NewsAll/:id", handlers.DeleteNewsById)

	r.Run(":8080")

}