package data

import (
	"clean/features/book"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) book.BookModel {
	return &model{
		connection: db,
	}
}

func (tm *model) InsertBook(pemilik string, inputBook book.Book) (book.Book, error) {
	var inputProcess = book.Book{
		Judul:   inputBook.Judul,
		Penulis: inputBook.Penulis,
		Genre:   inputBook.Genre,
		Tahun:   inputBook.Tahun,
		Pemilik: pemilik,
	}
	if err := tm.connection.Create(&inputProcess).Error; err != nil {
		return book.Book{}, err
	}

	return inputProcess, nil
}

func (tm *model) Update(pemilik string, todoID uint, data book.Book) (book.Book, error) {
	var qry = tm.connection.Where("pemilik = ? AND id = ?", pemilik, todoID).Updates(data)
	if err := qry.Error; err != nil {
		return book.Book{}, err
	}

	if qry.RowsAffected == 0 {
		return book.Book{}, errors.New("tidak ada data yg di update")
	}

	return data, nil
}

func (tm *model) GetBookByOwner(pemilik string) ([]book.Book, error) {
	var result []book.Book
	if err := tm.connection.Where("pemilik = ?  ", pemilik).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (m *model) DeleteBook(deleteID book.Book) error {
	var data = m.connection.Delete(&Book{}, deleteID)
	if err := data.Error; err != nil {
		return err
	}
	if data.RowsAffected == 0 {
		return errors.New("tidak ada data yg dihapus")
	}
	return nil
}
