package data

import "clean/features/book/data"

type User struct {
	Nama     string      `json:"judul" form:"judul" validate:"required,max=30,min=2"`
	Hp       string      `gorm:"type:varchar(13);primaryKey"`
	Password string      `json:"password" form:"password" validate:"required,max=30,min=8"`
	Books    []data.Book `gorm:"foreignKey:Pemilik;references:Hp"`
}
