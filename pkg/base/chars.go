// generate random string
// generate string md5 hash
package atomgo

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
	"unsafe"
)

const (
	// [a-z]+ lower case
	charsetLow = "abcdefghijklmnopqrstuvwxyz"
	// [A-Z]+ upper case
	charsetUp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// [0-9]+ digits
	charsetNum = "0123456789"
	// Character Normal
	charsetNormal = charsetLow + charsetUp + charsetNum
	// A few Special Characters in Common
	charsetSpecial = "@#$%"
	// All Special Characters include not common
	charsetExtra = charsetSpecial + "~!^&*()_+[]{}\\|;':\",./<>?"
	// base62 characters
	charsetBase62 = charsetNormal
	// base64 characters
	charsetBase64 = charsetNormal + "+/"
	// Hex characters
	charsetHex = "0123456789abcdef"
)

const (
	seedIdxBits = 6                  // 6 bits to represent a letter index
	seedIdxMask = 1<<seedIdxBits - 1 // All 1-bits, as many as letterIdxBits
	seedIdxMax  = 63 / seedIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func strGenerate(seedString string, n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), seedIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), seedIdxMax
		}
		if idx := int(cache & seedIdxMask); idx < len(seedString) {
			b[i] = seedString[idx]
			i--
		}
		cache >>= seedIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// generate random string from custom string
func StrRnd(seedStr string, length int) string {
	return strGenerate(seedStr, length)
}

// generate random string for characters
func StrRndChar(length int) string {
	return strGenerate(charsetLow+charsetUp, length)
}

// generate random string for number characters
func StrRndNum(length int) string {
	return strGenerate(charsetNum, length)
}

// generate random string for lower case characters include numbers
func StrRndLow(length int) string {
	return strGenerate(charsetLow+charsetNum, length)
}

// generate random string for upper case characters include numbers
func StrRndUp(length int) string {
	return strGenerate(charsetUp+charsetNum, length)
}

// generate random string for normal string
func StrRndNormal(length int) string {
	return strGenerate(charsetNormal, length)
}

// generate random string include special string
func StrRndSpecial(length int) string {
	return strGenerate(charsetNormal+charsetSpecial, length)
}

// generate random string include extra special string
func StrRndExtra(length int) string {
	return strGenerate(charsetNormal+charsetExtra, length)
}

// generate random string for base62 string
func StrRndBase62(length int) string {
	return strGenerate(charsetBase62, length)
}

// generate random string for base64 string
func StrRndBase64(length int) string {
	return strGenerate(charsetBase64, length)
}

// generate random string for hex string
func StrRndHex(length int) string {
	return strGenerate(charsetHex, length)
}

// generate md5 hash from string
func StrMd5(seedStr string) string {
	seedBytes := []byte(seedStr)
	b := md5.Sum(seedBytes)
	return hex.EncodeToString(b[:])
}
