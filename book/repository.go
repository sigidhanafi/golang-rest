package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
	Update(ID int, book Book) (Book, error)
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var dataBook Book
	err := r.db.Find(&dataBook, ID).Error

	return dataBook, err
}

func (r *repository) DeleteByID(ID int) error {
	err := r.db.Delete(&Book{}, ID).Error

	return err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(ID int, book Book) (Book, error) {
	var savedBook Book
	err := r.db.Where("id = ?", ID).First(&savedBook).Error

	if err != nil {
		return book, err
	}

	// update the data
	savedBook.Title = book.Title
	savedBook.Price = book.Price
	savedBook.Description = book.Description

	err = r.db.Save(&savedBook).Error

	if err != nil {
		return book, err
	}

	return savedBook, err
}
