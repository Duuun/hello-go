package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1234"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Yeah!!")
}
