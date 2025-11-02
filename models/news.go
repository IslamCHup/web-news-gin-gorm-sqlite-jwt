package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments" gorm:"constraint:OnDelete:CASCADE; foreignKey:NewsID"`       
}
