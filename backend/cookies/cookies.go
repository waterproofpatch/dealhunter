package cookies

import "net/http"

func SetRefreshTokenCookie(refreshToken string) *http.Cookie {
	// Create a cookie
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	return cookie
}
