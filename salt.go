package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|"
	letterIdxBits = 7                    // 7 bits to represent a letter index (92 is higher than 64 (2<<6) but lower than 128 (2<<7))
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	saltTypes   = [8]string{"AUTH_KEY", "SECURE_AUTH_KEY", "LOGGED_IN_KEY", "NONCE_KEY", "AUTH_SALT", "SECURE_AUTH_SALT", "LOGGED_IN_SALT", "NONCE_SALT"}
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|")
	src         = rand.NewSource(time.Now().UnixNano())
)

func main() {
	http.Handle("/", get(GiveSalts))
	http.Handle("/env", get(GiveSaltsEnv))
	http.Handle("/json", get(GiveSaltsJSON))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8090"
	}

	log.Printf("Starting to listen on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func get(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		f(w, r)
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GiveSalts responds to the GET / request
func GiveSalts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(GenerateSaltsWP512()))
}

// GiveSaltsEnv responds to the GET /env request
func GiveSaltsEnv(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(GenerateSaltsEnv512()))
}

// GiveSaltsJSON responds to the GET /json request
func GiveSaltsJSON(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(GenerateSaltsJSON512())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal JSON"))
		log.Printf("failed to marshal json: %v", err)
	}

	w.Write(resp)
}

// GenerateSaltsWP512 generates the content for GiveSalts by calling the method once and slicing
func GenerateSaltsWP512() string {
	formattedStrings := make([]string, 8)
	longstring := RandStringBytesMaskImpr(512)

	for i, arg := range saltTypes {
		formattedStrings[i] = fmt.Sprintf("define( '%s',%s'%s' );", arg, strings.Repeat(" ", 17-len(arg)), longstring[i*64:(i+1)*64])
	}
	return strings.Join(formattedStrings, "\n")
}

// GenerateSaltsEnv512 generates the content for GiveSaltsEnv by calling the method once and slicing
func GenerateSaltsEnv512() string {
	formattedStrings := make([]string, 8)
	longstring := RandStringBytesMaskImpr(512)

	for i, arg := range saltTypes {
		formattedStrings[i] = fmt.Sprintf("%s=\"%s\"", arg, longstring[i*64:(i+1)*64])
	}
	return strings.Join(formattedStrings, "\n")
}

// GenerateSaltsJSON512 generates the content for GiveSaltsJSON by calling the method once and slicing
func GenerateSaltsJSON512() map[string]string {
	formattedStrings := make(map[string]string)
	longstring := RandStringBytesMaskImpr(512)

	for i, arg := range saltTypes {
		formattedStrings[arg] = longstring[i*64 : (i+1)*64]
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
