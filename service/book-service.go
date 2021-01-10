package service

import (
	"github.com/roblesok/gorm-fiber-crud/entity"
	"github.com/roblesok/gorm-fiber-crud/repository"
)

type BookService interface {
	Insert(book *entity.Book) (*entity.Book, error)
	FetchAll() (*[]entity.Book, error)
	FetchOne(ID int) (*entity.Book, error)
	Delete(ID int) error
}

type service struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &service{repo: r}
}

func (s *service) Insert(book *entity.Book) (*entity.Book, error) {
	return s.repo.Add(book)
}

func (s *service) FetchAll() (*[]entity.Book, error) {
	return s.repo.FindAll()
}

func (s *service) FetchOne(ID int) (*entity.Book, error) {
	return s.repo.FindOne(ID)
}

func (s *service) Delete(ID int) error {
	return s.repo.Delete(ID)
}
