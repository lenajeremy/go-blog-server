package models

import "github.com/google/uuid"

type User struct {
	BaseModel
	Posts     []Post
	ProfileID uuid.UUID `json:"profileID" gorm:"profile_id;not null;type:uuid"`
	Profile   Profile   `json:"profile" gorm:"not null;foreignKey=ProfileID;references=ID"`
}
