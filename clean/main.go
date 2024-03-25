package main

import (
	"clean/config"
	td "clean/features/book/data"
	th "clean/features/book/handler"
	ts "clean/features/book/services"
	"clean/features/user/data"
	"clean/features/user/handler"
	"clean/features/user/services"
	"clean/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	bookData := td.New(db)
	bookService := ts.NewBookService(bookData)
	bookHandler := th.NewHandler(bookService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, userHandler, bookHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
