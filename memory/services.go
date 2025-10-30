package memory

import (
	"github.com/IslamCHup/web-news-service-gin/models"
)

var NewsAll = []models.News{
	{
		Id:      0,
		Title:   "В России запртеили",
		Text:    "Запретили запрещенное...",
		Picture: "https://sibadvokat.ru/wp-content/uploads/2023/10/sudebnyj-zapret-na-priblizhenie-predlozheno-vvesti-v-rossii-730x430.jpg",
	},
	{
		Id:      1,
		Title:   "Собка",
		Text:    "Красивое собака...",
		Picture: "https://images.wagwalkingweb.com/media/daily_wag/blog_articles/hero/1723114015.705158/popular-dogs-hero-1.jpg",
	},
}

var NextInd = 2

func GetAll()[]models.News{
	return NewsAll
}

func GetNewsId(id int)models.News{
	for i := range NewsAll{
		if NewsAll[i].Id == id {
			return NewsAll[i]
		}
	}
	return models.News{}
}

func Create(n models.News) models.News{
	n.Id = NextInd
	NextInd++
	NewsAll = append(NewsAll, n)
	return n
}

func Update(id int, n models.News)bool{
	for i := range NewsAll{
		if NewsAll[i].Id == id{
			NewsAll[i] = n
			return true
		}
	}
	return false
}

func Delete(id int)bool{
	for i := range NewsAll{
		if NewsAll[i].Id == id{
			NewsAll = append(NewsAll[:i], NewsAll[i+1:]...)
			return true
		}
	}
	return false
}