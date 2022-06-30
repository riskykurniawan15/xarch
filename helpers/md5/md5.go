package md5

import (
	"crypto/md5"
	"fmt"
)

func Encrypt(value string) string {
	data := []byte(value)
	return fmt.Sprintf("%x", md5.Sum(data))
}
