package Services

import (
	"context"
	model "go-test/Model"
	repository "go-test/Repository"
)

type IBookService interface {
	GetAllBooks(ctx context.Context) ([]model.Book, error)
}

type BookService struct {
	bookRepo repository.IRepository
}

func NewBookService(bookRepo repository.IRepository) *BookService {
	return &BookService{bookRepo: bookRepo}
}

func (u *BookService) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var books []model.Book
	err := u.bookRepo.GetAll(ctx, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
