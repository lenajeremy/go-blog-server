package models

type Tag struct {
	BaseModel
	Text  string `json:"text" gorm:"not null"`
	Posts []Post `json:"posts" gorm:"not null"`
}
