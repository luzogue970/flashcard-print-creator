package models

type Flashcard struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"title"`
	Front    string `json:"front"`  // Question/Description
	Back     string `json:"back"`   // Réponse
}