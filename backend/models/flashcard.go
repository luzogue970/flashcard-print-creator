package models

import "gorm.io/gorm"

type Flashcard struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey" json:"id"`
	CategoryID        uint   `json:"category_id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	InformationNumber int    `json:"information_number"`
	Answers           string `json:"answers"`
}

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
