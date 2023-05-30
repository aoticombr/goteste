package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func StrToMd5(value string) string {

	hash := md5.Sum([]byte(value))            // Calcula o hash MD5
	hashString := hex.EncodeToString(hash[:]) // Converte o hash para uma string hexadecimal

	return hashString

}
