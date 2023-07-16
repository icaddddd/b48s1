package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"name":   "Rhisjad",
			"addres": "jalan margonda",
		})
	})

	e.GET("/about", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "HELLO WORLD",
		})
	})

	e.Logger.Fatal(e.Start("localhost:5000"))
}
