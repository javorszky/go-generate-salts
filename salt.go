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
const (
	letterIdxBits = 7                    // 7 bits to represent a letter index (92 is higher than 64 (2<<6) but lower than 128 (2<<7))
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_ []{}<>~`+=,.;:/?|")
var src = rand.NewSource(time.Now().UnixNano())

func main() {
	e := echo.New()

	e.GET("/", giveSalts)
	e.GET("/env", giveSaltsEnv)

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

func giveSaltsEnv(c echo.Context) error {
	string := RandStringBytesMaskImprSrc(512)
	return c.String(http.StatusOK, string)
}

// define('AUTH_KEY',         'H!R+>kY|$Nf|Nx%ElcRs; 1?;7[JH63-`F!)jVJ<,&jLY,{f+spcv+r+hRR6tVrA');
// define('SECURE_AUTH_KEY',  'qWtn8!1pN.20WMoh_U-yCPu>.fi@PfIMSq2`mR!#H[32m8QT-q|R<O%,i+e~T+yr');
// define('LOGGED_IN_KEY',    'M6~Cg)qsd}ke(nlb= ;0Qu:+b:+C6n&30ngvGq-k!yj^Xs>{XBwpj-x!brRmR??F');
// define('NONCE_KEY',        'Jx]|.<)WzY4grFWnlmGdI,`(v#cxTvoOtbrlV-$5]r:z|J)9Ouc,Go .YBW{v^nk');
// define('AUTH_SALT',        'L^D-sxa}@iu{*r3i+|qm;,-l/7Po+S?P<=iCMqh;,EJ/7D8`[1yj]: vkLF-^:f1');
// define('SECURE_AUTH_SALT', '`l_vCu1htqxQ+i8+15$XfFuBzj65l!pAT=x^|PL8,Le%Kq`[wY)km e-p|4r~rws');
// define('LOGGED_IN_SALT',   '5&nNK-aXkIhFD01b&EJ#L>JeW3]%{((vBtnst.`X_~<w&lO+bpQtvW@A7GSu(dA,');
// define('NONCE_SALT',       'wwsQ#<zh3+y> 1o`VH,db>%MD7|qG-:c_f!#.OuBp|)<%rRz.jscc[}=q>T5-TpZ');

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

// RandStringBytesRmndr generates random string from the remainder from dividing a 64 bit number with the 64 bit representation of the length of characters
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// RandStringBytesMask only uses the last 7 bits of the generated 63 byte int
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
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

// RandStringBytesMaskImprSrc slices the rand number and uses a rand.Source instead of a rand.Rand
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
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
