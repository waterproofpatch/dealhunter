package cookies

import (
	"net/http"
	"time"
)

func SetRefreshTokenCookie(refreshToken string) *http.Cookie {
	// Create a cookie
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * time.Hour), // The cookie will expire in 24 hours
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	return cookie
}
