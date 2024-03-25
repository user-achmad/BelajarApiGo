package routes

import (
	"clean/config"
	book "clean/features/book"
	user "clean/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, tc book.BookController) {
	userRoute(c, ctl)
	bookRoute(c, tc)
}

func userRoute(c *echo.Echo, ctl user.UserController) {
	// input untuk user
	c.POST("/users", ctl.Add())
	c.POST("/login", ctl.Login())
	c.GET("/users", ctl.View(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/users/:hp", ctl.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/users/:hp", ctl.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func bookRoute(c *echo.Echo, tc book.BookController) {
	// input untuk buku
	c.GET("/book", tc.View(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.POST("/book", tc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/book/:id", tc.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/book/:id", tc.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
