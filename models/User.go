package models

type User struct {
	BaseModel
	Posts         []Post    `json:"posts" gorm:"posts;foreignKey:AuthorID"`
	Comments      []Comment `json:"comments" gorm:"comments;foreignKey:AuthorID"`
	Email         string    `json:"email" gorm:"email;unique;not null"`
	Password      string    `json:"password" gorm:"password;not null"`
	EmailVerified bool      `json:"emailVerified" gorm:"email_verified;default:false;not null"`
}

type PublicUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
	ID        string `json:"ID"`
}
