package shouqianba

import (
	"crypto/md5"
	"encoding/hex"
)

func getSign(signStr string) string {
	md5Str, _ := MD5(signStr)
	return md5Str
}

// MD5 加密
func MD5(str string) (string, error) {
	hs := md5.New()
	if _, err := hs.Write([]byte(str)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hs.Sum(nil)), nil
}
