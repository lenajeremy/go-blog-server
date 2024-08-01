package models

import "github.com/google/uuid"

type Profile struct {
	BaseModel
	FirstName    *string `json:"firstName" gorm:"first_name"`
	LastName     *string `json:"lastName" gorm:"last_name"`
	Age          *uint   `json:"age" gorm:"age"`
	DefaultTheme string  `json:"defaultTheme" gorm:"default_theme"`

	UserId uuid.UUID `json:"userId" gorm:"user_id;not null"`
	User   User      `json:"user" gorm:"foreignKey:UserId"`
}
