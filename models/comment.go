package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	AuthorId    int  `json:"author_id"`
	Author      User `gorm:"foreignKey:AuthorId"`
	NewsId      int	 `json:"news_id"`
	News        News `json:"news" gorm:"foreignKey:NewsId"`
	TextComment string `json:"text_comment"`
}
