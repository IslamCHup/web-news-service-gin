package handlers

import (
	"net/http"
	"strconv"

	"github.com/IslamCHup/web-news-service-gin/memory"
	"github.com/IslamCHup/web-news-service-gin/models"
	"github.com/gin-gonic/gin"
)



func GetAllNews(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, memory.GetAll())
}

func GetNewsById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	news := memory.GetNewsId(id)
	c.IndentedJSON(http.StatusOK, news)
}

func CreateNews(c *gin.Context) {
	create := models.News{}
	err := c.ShouldBindJSON(&create)
	if err != nil {
		c.String(http.StatusBadGateway, "Неправильный запрос")
		return
	}
	created := memory.Create(create)
	c.JSON(http.StatusOK, created)
}

func UpdateNews(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	needUpd := models.News{}
	err := c.ShouldBindJSON(&needUpd)
	if err != nil{
		c.String(http.StatusBadRequest, "Нет такой новости")
		return
	}
	isUpdated := memory.Update(id, needUpd)
	wordUpd := ""
	if isUpdated {
		wordUpd = "новость обновлена"
	} else {
		wordUpd = "ошибка, новость не обновлена"
	}
	c.String(http.StatusOK, "Результаты обновления: %s", wordUpd)
}

func DeleteNewsById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	isDeleted := memory.Delete(id)
	wordDel := ""
	if isDeleted{
		wordDel = "новость удалена"
	} else {
		wordDel = "новость не удалена"
	}
	c.String(http.StatusOK, "Результаты удаления: %s", wordDel)

}
