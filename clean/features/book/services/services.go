package services

import (
	"clean/features/book"
	"clean/helper"
	"clean/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m book.BookModel
	v *validator.Validate
}

func NewBookService(model book.BookModel) book.BookService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) View(token *jwt.Token) ([]book.Book, error) {
	decodeHp := middlewares.DecodeToken(token)
	result, err := s.m.GetBookByOwner(decodeHp)
	if err != nil {
		return []book.Book{}, err
	}

	return result, nil
}

func (s *service) AddBook(pemilik *jwt.Token, kegiatanBaru book.Book) (book.Book, error) {
	hp := middlewares.DecodeToken(pemilik)
	if hp == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return book.Book{}, errors.New("data tidak valid")
	}

	err := s.v.Struct(&kegiatanBaru)
	if err != nil {
		log.Println("error validasi", err.Error())
		return book.Book{}, err
	}

	result, err := s.m.InsertBook(hp, kegiatanBaru)
	if err != nil {
		return book.Book{}, errors.New(helper.ServerGeneralError)
	}

	return result, nil
}

func (s *service) Update(hp string, todoID uint, newData book.Book) error {
	var updateValidate book.Update
	updateValidate.Judul = newData.Judul
	updateValidate.Penulis = newData.Penulis
	updateValidate.Genre = newData.Genre
	updateValidate.Tahun = newData.Tahun
	err := s.v.Struct(&updateValidate)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	_, err = s.m.Update(hp, todoID, newData)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}

func (s *service) DeleteBook(deleteID book.Book) error {
	var deleteUser book.DeleteBook
	deleteUser.ID = deleteID.ID
	err := s.v.Struct(&deleteUser)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}
	err = s.m.DeleteBook(deleteID)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}
	return nil
}
