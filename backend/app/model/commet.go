package model

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment  string    `json:"password"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}