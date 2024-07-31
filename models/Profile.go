package models

type Profile struct {
	FirstName    *string `json:"firstName" gorm:"first_name"`
	LastName     *string `json:"lastName" gorm:"last_name"`
	Age          *uint   `json:"age" gorm:"age"`
	DefaultTheme string  `json:"defaultTheme" gorm:"default_theme"`
}
