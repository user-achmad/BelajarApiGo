package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// bagian yang berisi KONTRAK mengenai obyek yang digunakan / disepakati dalam proses coding kalian

type UserController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Login() echo.HandlerFunc
	View() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	// CekUser(hp string) bool
	Update(hp string, data User) error
	Delete(hp User) error
	Login(loginData User) (User, string, error)
	View(token *jwt.Token) (User, error)
	// Delete()
}

type UserModel interface {
	InsertUser(newData User) error
	DeleteUser(newData User) error
	Update(hp string, data User) error
	Login(hp string) (User, error)
	GetUserByHP(hp string) (User, error)
}

type User struct {
	Nama     string
	Hp       string
	Password string
}

type Login struct {
	Hp       string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}

type Register struct {
	Nama     string `validate:"required"`
	Hp       string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}

type Update struct {
	Nama     string `validate:"required"`
	Hp       string `validate:"required,min=10,max=13,numeric"`
	Password string `validate:"required,alphanum,min=8"`
}

type Delete struct {
	Hp string `validate:"required"`
}
