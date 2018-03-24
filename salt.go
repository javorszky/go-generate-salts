package main

import (
	"fmt"
	"net/http"
	"os"
	s "strings"

	"github.com/labstack/echo"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|"

func main() {
	e := echo.New()

	e.GET("/", giveSalts)

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Print("Port not in env, setting it to 8090")
		port = "8090"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func giveSalts(c echo.Context) error {

	return c.String(http.StatusOK, "Hellow, World"+s.Repeat("test", 5))
}
