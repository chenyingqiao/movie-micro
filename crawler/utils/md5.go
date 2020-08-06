package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//Md5Simple md5
func Md5Simple(str string) string {
	var buf []byte
	help := md5.New()
	help.Write(buf)
	return hex.EncodeToString(help.Sum(nil))
}
