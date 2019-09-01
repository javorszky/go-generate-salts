package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	echopprof "github.com/sevenNt/echo-pprof"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#%^&*()-_ []{}<>~`+=,.;:/?|"
	letterIdxBits = 7                    // 7 bits to represent a letter index (92 is higher than 64 (2<<6) but lower than 128 (2<<7))
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	saltTypes   = [9]string{"AUTH_KEY", "SECURE_AUTH_KEY", "LOGGED_IN_KEY", "NONCE_KEY", "AUTH_SALT", "SECURE_AUTH_SALT", "LOGGED_IN_SALT", "NONCE_SALT", "WP_CACHE_KEY_SALT"}
	saltTypeLen = len(saltTypes)
	saltBytes   = saltTypeLen * 64
)

func main() {
	e := echo.New()

	e.GET("/", GiveSalts)
	e.GET("/env", GiveSaltsEnv)
	e.GET("/json", GiveSaltsJSON)

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Print("Port not in env, setting it to 8090")
		port = "8090"
	}

	echopprof.Wrap(e)

	e.Logger.Fatal(e.Start(":" + port))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GiveSalts responds to the GET / request
func GiveSalts(c echo.Context) error {
	return c.String(http.StatusOK, GenerateSaltsWPEfficient())
}

// GiveSaltsEnv responds to the GET /env request
func GiveSaltsEnv(c echo.Context) error {
	return c.String(http.StatusOK, GenerateSaltsEnvEfficient())
}

// GiveSaltsJSON responds to the GET /json request
func GiveSaltsJSON(c echo.Context) error {
	return c.JSON(http.StatusOK, GenerateSaltsJSONEfficient())
}

// GenerateSaltsWPEfficient generates the content for GiveSalts by calling the method once and slicing
func GenerateSaltsWPEfficient() string {
	formattedStrings := make([]string, saltTypeLen)
	longString := RandStringBytesMaskImpr(saltBytes)

	for i, arg := range saltTypes {
		formattedStrings[i] = fmt.Sprintf("define( '%s',%s'%s' );", arg, strings.Repeat(" ", 18-len(arg)), longString[i*64:(i+1)*64])
	}
	return strings.Join(formattedStrings, "\n")
}

// GenerateSaltsEnvEfficient generates the content for GiveSaltsEnv by calling the method once and slicing
func GenerateSaltsEnvEfficient() string {
	formattedStrings := make([]string, saltTypeLen)
	longString := RandStringBytesMaskImpr(saltBytes)

	for i, arg := range saltTypes {
		formattedStrings[i] = fmt.Sprintf("%s=\"%s\"", arg, longString[i*64:(i+1)*64])
	}
	return strings.Join(formattedStrings, "\n")
}

// GenerateSaltsJSONEfficient generates the content for GiveSaltsJSON by calling the method once and slicing
func GenerateSaltsJSONEfficient() map[string]string {
	formattedStrings := make(map[string]string)
	longString := RandStringBytesMaskImpr(saltBytes)

	for i, arg := range saltTypes {
		formattedStrings[arg] = longString[i*64 : (i+1)*64]
	}
	return formattedStrings
}

// RandStringBytesMaskImpr slices up the bits of the random number and uses all slices
func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
