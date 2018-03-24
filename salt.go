package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	s "strings"
	"time"

	"github.com/labstack/echo"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|")

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func giveSalts(c echo.Context) error {
	return c.String(http.StatusOK, "Hellow, World"+s.Repeat("test", 5))
}

// RandStringRunes generates random string runes
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandStringBytes generates random string from bytes
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
