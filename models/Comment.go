package models

import "github.com/google/uuid"

type Comment struct {
	BaseModel
	Content  string    `json:"content" gorm:"not null"`
	AuthorID uuid.UUID `json:"authorID" gorm:"author_id;type:uuid;not null"`
	//Replies          []Comment `json:"replies"`
	PostCommentedOn uuid.UUID `json:"postId" gorm:"post_commented_on;type:uuid;not null;"`
	//CommentRepliedTO uuid.UUID `json:"parentCommentId" gorm:"comment_replied_to;type:uuid;null"`
}
