package utils

import "crypto/md5"

func GenerateToken(usuario string, senha string) string {
	encoder := md5.New()
	encoder.Write([]byte(usuario + senha))
	return string(encoder.Sum(nil))
}
