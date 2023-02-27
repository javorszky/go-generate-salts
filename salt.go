package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#%^&$*()-_ []{}<>~`+=,.;:/?|"
	letterIdxBits = 7                    // 7 bits to represent a letter index (92 is higher than 64 (2<<6) but lower than 128 (2<<7))
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	saltTypes = [9]string{
		"AUTH_KEY",
		"SECURE_AUTH_KEY",
		"LOGGED_IN_KEY",
		"NONCE_KEY",
		"AUTH_SALT",
		"SECURE_AUTH_SALT",
		"LOGGED_IN_SALT",
		"NONCE_SALT",
		"WP_CACHE_KEY_SALT",
	} // all the different salts we're going to generate
	saltTypeLen = len(saltTypes)   // how many of them are there
	saltBytes   = saltTypeLen * 64 // total number of bytes we need to generate a long string that we're going to cut to size.
)

func main() {
	e := echo.New()

	e.GET("/", giveSalts)
	e.GET("/env", giveSaltsEnv)
	e.GET("/json", giveSaltsJSON)
	e.GET("/docker", giveSaltsDocker)

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Print("Port not in env, setting it to 8090")

		port = "8090"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

// giveSalts responds to the GET / request
func giveSalts(c echo.Context) error {
	return c.String(http.StatusOK, generateSaltsWPEfficient())
}

// giveSaltsEnv responds to the GET /env request
func giveSaltsEnv(c echo.Context) error {
	return c.String(http.StatusOK, generateSaltsEnvEfficient())
}

// giveSaltsJSON responds to the GET /json request
func giveSaltsJSON(c echo.Context) error {
	return c.JSON(http.StatusOK, generateSaltsJSONEfficient())
}

func giveSaltsDocker(c echo.Context) error {
	rnd := make([]byte, 1)
	n, err := crand.Read(rnd)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("%d bytes, and it's\n\n"+
		"%08b\n"+
		"%08b\n"+
		"%08b", n, rnd[0], letterIdxMax, rnd[0]&letterIdxMax))
}

// generateSaltsWPEfficient generates the content for giveSalts by calling the method once and slicing
func generateSaltsWPEfficient() string {
	formattedStrings := make([]string, saltTypeLen)
	longString := randStringBytesMaskImpr(saltBytes)

	for i, arg := range saltTypes {
		formattedStrings[i] = fmt.Sprintf("define( '%s',%s'%s' );", arg, strings.Repeat(" ", 18-len(arg)), longString[i*64:(i+1)*64])
	}

	return strings.Join(formattedStrings, "\n")
}

// generateSaltsEnvEfficient generates the content for giveSaltsEnv by calling the method once and slicing
func generateSaltsEnvEfficient() string {
	formattedStrings := make([]string, saltTypeLen)
	longString := randStringBytesMaskImpr(saltBytes)

	for i, arg := range saltTypes {
		formattedStrings[i] = fmt.Sprintf("%s=\"%s\"", arg, longString[i*64:(i+1)*64])
	}

	return strings.Join(formattedStrings, "\n")
}

// generateSaltsJSONEfficient generates the content for giveSaltsJSON by calling the method once and slicing
func generateSaltsJSONEfficient() map[string]string {
	formattedStrings := make(map[string]string)
	longString := randStringBytesMaskImpr(saltBytes)

	for i, arg := range saltTypes {
		formattedStrings[arg] = longString[i*64 : (i+1)*64]
	}

	return formattedStrings
}

// randStringBytesMaskImpr slices up the bits of the random number and uses all slices
func randStringBytesMaskImpr(n int) string {
	b := make([]byte, n)

	//rnd := make([]byte, 8)

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
