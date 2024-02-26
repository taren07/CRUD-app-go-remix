package usecase

import (
	"app/model"
)

type ICommentUsecase interface {
	CreateComment(comment model.Comment) (model.CommentResponse, error)
}

type commentUsecase struct {
	// cr repository.ICommentRepository
}

