package repository

import (
	"app/model"
)

type ICommentRepository interface {
	GetAllComments(comments *[]model.Comment, userId uint) error
}