package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// "web-news-gin-gorm-sqlite-jwt/database"
	"web-news-gin-gorm-sqlite-jwt/models"
)




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




