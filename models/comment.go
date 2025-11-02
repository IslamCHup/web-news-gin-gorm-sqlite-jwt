package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	AuthorID    int    `json:"author_ID"`
	Author      User   `gorm:"foreignKey:AuthorID"`
	NewsID      int    `json:"news_ID"`
	News        News   `json:"news" gorm:"foreignKey:NewsID"`
	TextComment string `json:"text_comment"`
}
