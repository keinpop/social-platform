package middleware

import (
	"mai-platform/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

const appKey = "app"

func WithApp(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(appKey, app)

		c.Next()
	}
}

func GetApp(c *gin.Context) *app.App {
	value, ok := c.Get(appKey)
	if !ok {
		return nil
	}

	a, ok := value.(*app.App)
	if !ok {
		return nil
	}

	return a
}

func WithAuth(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, err := app.Auth.CheckToken(c.Request.Header.Get("Authorization"))
		if !ok || err != nil {
			c.JSON(http.StatusUnauthorized, "")
			c.Abort()
			return
		}

		c.Next()
	}
}

// func WithAuth - из контекста достает header авторизации Authorization и отправлять его значение
// в api-авторизации в таком же хеддере, если сервис авторизации возвращает 200
// тогда middleware продолжает запрос
// иначе отдает пользователю ответ, который получила
//  curl localhost:8090/auth/hello -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA0ODIyMDksImlkIjoiYWJjNEBtYWlsLnJ1Iiwib3JpZ19pYXQiOjE3MjA0Nzg2MDl9.0Um3f24ranhbp3wy83Q2pS4iPk_vwhXmQ6O05L8rpKg"
//
