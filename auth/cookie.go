package auth

import (
	"net/http"
	"time"
)

// NewCookie creates a new auth cookie with an expiry length of 1 year
func NewCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:    "remember_token",
		Value:   token,
		Expires: time.Now().AddDate(1, 0, 0), // 1 Year
	}
}

// ExpireCookie creates a new cookie with Expires set to now
func ExpireCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "remember_token",
		Value:   "",
		Expires: time.Now(),
	}
}
