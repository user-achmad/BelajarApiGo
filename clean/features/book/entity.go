package book

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookController interface {
	View() echo.HandlerFunc
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type BookModel interface {
	GetBookByOwner(pemilik string) ([]Book, error)
	InsertBook(pemilik string, kegiatanBaru Book) (Book, error)
	Update(pemilik string, todoID uint, data Book) (Book, error)
	DeleteBook(DeleteBook Book) error
}

type BookService interface {
	View(token *jwt.Token) ([]Book, error)
	Update(pemilik string, todoID uint, inputBook Book) error
	AddBook(pemilik *jwt.Token, kegiatanBaru Book) (Book, error)
	DeleteBook(deleteID Book) error
	// UpdateTodo(pemilik *jwt.Token, todoID string, data Todo) (Todo, error)
}

type Book struct {
	gorm.Model
	Judul   string
	Penulis string
	Genre   string
	Tahun   string
	Pemilik string
}
type BookRequest struct {
	Judul   string
	Penulis string
	Genre   string
	Tahun   string
	Pemilik string
}
type Update struct {
	gorm.Model
	Judul   string
	Penulis string
	Genre   string
	Tahun   string
}

type DeleteBook struct {
	gorm.Model
}
