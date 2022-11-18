package passwd

import (
	"golang.org/x/crypto/bcrypt"
)

// Check 检查密码的正确性
func Check(password, hash string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
