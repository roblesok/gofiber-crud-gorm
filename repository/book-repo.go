package repository

import "github.com/roblesok/gorm-fiber-crud/entity"

type BookRepository interface {
	Add(post *entity.Book) (*entity.Book, error)
	FindAll() (*[]entity.Book, error)
	FindOne(ID int) (*entity.Book, error)
	Delete(ID int) error
}
