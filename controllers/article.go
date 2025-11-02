package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"web-news-gin-gorm-sqlite-jwt/database"
	"web-news-gin-gorm-sqlite-jwt/models"
)

/*
 controllers/ — логика обработки запросов
Зачем: Это мозг API. Здесь ты определяешь, что происходит, когда приходит запрос.
Три файла:
auth.go — регистрация и логин пользователей.
article.go — операции со статьями (создать, получить список, удалить, и т.д.).
comment.go — работа с комментариями.
Каждый контроллер:
Получает входные данные из запроса.
Проверяет их.
Работает с базой через database.DB.
Возвращает JSON-ответ.
💡 Подумай о контроллере как о “мини-функции с бизнес-логикой для одного запроса”.
*/


func GettAllNews(db *gorm.DB)  gin.HandlerFunc{
	return func (c *gin.Context){
		var news []models.News

		db.Preload("Comments.Author").Find(&news)
		c.JSON(http.StatusOK, news)
	}
}

func GetNewsById(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		var news models.News

		if err := db.Preload("Comments.Author", ).Where("id = ?", c.Param("id")).First(&news).Error; err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		}
	
		c.JSON(http.StatusOK, news)
	}
}



