package cookie

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func CreateCookie(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateCookieLogin(cookie string) string {
	cookie_new := fmt.Sprintf("auth_cookie=%s; Max-Age=%d; path=/", cookie, 7*24*3600)
	return string(cookie_new)
}

func CreateCookieLogout() string {
	cookie_new := fmt.Sprintf("auth_cookie=none; Max-Age=0; path=/")
	return string(cookie_new)
}
