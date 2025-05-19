package password

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Password(username, password string) string {
	return Sha256(fmt.Sprintf("%s-%s", username, password))
}

func Verify(username, newPass, oldPass string) bool {
	return Password(username, newPass) == oldPass
}
