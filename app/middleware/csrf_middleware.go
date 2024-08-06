package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
)

// CSRFMiddleware создает middleware для CSRF защиты
func CSRFMiddleware() gin.HandlerFunc {
	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false)) // Установите Secure(true) для HTTPS
	return func(c *gin.Context) {
		// Проверка CSRF токена
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})
		CSRF(handler).ServeHTTP(c.Writer, c.Request)

		// Добавление CSRF токена в контекст
		csrfToken := csrf.Token(c.Request)
		c.Set("CSRFToken", csrfToken)
	}
}
