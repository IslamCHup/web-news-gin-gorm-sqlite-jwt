package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"web-news-gin-gorm-sqlite-jwt/controllers"
	"web-news-gin-gorm-sqlite-jwt/middleware"
)

/*
routes/routes.go — маршрутизация (роутинг)

Зачем: Здесь ты определяешь, какие URL-адреса обрабатывает твой сервер и какие функции вызываются.

В нём ты создаёшь экземпляр gin.Engine и описываешь:

Публичные маршруты (/register, /login, /articles).

Приватные маршруты (например, /auth/articles, которые требуют JWT).

💡 Тут не должно быть логики, только маршруты и подключение к контроллерам.
*/
func SetupGroup(r *gin.Engine, db *gorm.DB) {
	{
		groupNews := r.Group("/news")

		groupNews.GET("/", controllers.GettAllNews(db))
		groupNews.GET("/:id", controllers.GetNewsById(db))
	}
	{
		authNews := r.Group("/news", middleware.AuthReq())

		authNews.POST("/", controllers.CreateNews(db))
		authNews.PATCH("/:id", controllers.UpdateNews(db))
		authNews.DELETE("/:id", controllers.DeleteNews(db))
	}
	{
		groupComments := r.Group("/comment")

		groupComments.GET("/news/:newsId", controllers.GetCommentsOfNews(db))
		groupComments.GET("/:id", controllers.GetCommentsById(db))
		groupComments.GET("/author/:authorId", controllers.GetCommentAuthorById(db))
	}
	
	{
		authComments := r.Group("/comment", middleware.AuthReq())

		authComments.POST("/", controllers.CreateComment(db))
		authComments.DELETE("/:id", controllers.DeleteCommentsById(db))
	}
	{
		groupAuth := r.Group("/auth")

		groupAuth.POST("/register", controllers.Registration)
		groupAuth.POST("/login", controllers.Login)
	}
}
