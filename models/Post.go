package models

import "github.com/google/uuid"

type Post struct {
	BaseModel
	Title    string    `json:"title" gorm:"title; not null"`
	SubTitle string    `json:"subtitle" gorm:"subtitle"`
	Content  string    `json:"content" gorm:"content"`
	Comments []Comment `jsons:"comments" gorm:"foreignKey:PostCommentedOn"`
	AuthorID uuid.UUID `json:"authorID" gorm:"author_id;not null;type:uuid"`
}
