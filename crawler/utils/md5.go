package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Simple(str string) string {
	var buf []byte
	help := md5.New()
	help.Write(buf)
	return hex.EncodeToString(help.Sum(nil))
}
