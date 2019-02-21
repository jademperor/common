package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	// "strings"
)

var (
	lenOfUUIDBytes = 48
)

// StringMD5 ...
func StringMD5(s string) string {
	m := md5.New()
	io.WriteString(m, s)
	return hex.EncodeToString(m.Sum(nil))
}

// Fstring ...
func Fstring(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}

// MD5 md5 encrypt
func MD5(bs []byte) string {
	m := md5.New()
	m.Write(bs)
	return hex.EncodeToString(m.Sum(nil))
}

// UUID get a random UUID string without any parameter
func UUID() string {
	bs := make([]byte, lenOfUUIDBytes)
	if _, err := io.ReadFull(rand.Reader, bs); err != nil {
		println(err)
		return ""
	}
	// md5 := MD5(bs)
	// to upper
	// upper := strings.ToUpper(md5)
	return MD5(bs)[:6]
}

// SetUUIDBytesLen to control the UUID length
func SetUUIDBytesLen(n int) {
	if n <= 0 {
		return
	}
	lenOfUUIDBytes = n
}
