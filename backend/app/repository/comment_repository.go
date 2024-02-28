package repository

import (
	"app/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	GetAllComments(comments *[]model.Comment, userId uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}