package main

import (
	"github.com/labstack/echo"
	"net/http"
)

const helloMessage = "Hello, World!"

func main() {
	router := NewRouter()
	router.Start(":8080")
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/hello", helloHandler)
	return e
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, helloMessage)
}
