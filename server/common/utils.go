package common

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(msg string) string {
	m := md5.New()
	m.Write([]byte(msg))
	return hex.EncodeToString(m.Sum(nil))
}

func EncodePassword(origin string) string {
	return Md5(Md5(origin))
}
