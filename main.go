package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"web-news-gin-gorm-sqlite-jwt/database"
	"web-news-gin-gorm-sqlite-jwt/models"
)

/*
В нём ты:

Подключаешься к базе данных.

Делаешь миграции (чтобы таблицы создались).

Создаёшь роутер (из routes/).

Запускаешь сервер на порту (например, :8080).

*/

func main() {
	db, err := database.SetupDB()

	if err != nil{
		panic("Failed connection to DB" + err.Error())
	}

	r := gin.Default()

	//routes

	r.Run(":8080")
}
