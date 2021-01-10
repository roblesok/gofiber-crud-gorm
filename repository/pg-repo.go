package repository

import (
	"errors"

	"github.com/roblesok/gorm-fiber-crud/entity"
)

type repo struct{}

func NewPgRepository() BookRepository {
	return &repo{}
}

func (r *repo) Add(book *entity.Book) (*entity.Book, error) {
	return nil, errors.New("TODO Add BOOK")
}

func (r *repo) FindAll() (*[]entity.Book, error) {
	return nil, errors.New("TODO FindAllBoks")
}

func (r *repo) FindOne(ID int) (*entity.Book, error) {
	return nil, errors.New("TODO FindBook")
}

func (r *repo) Delete(postId int) error {
	return errors.New("TODO Delete book")
}
