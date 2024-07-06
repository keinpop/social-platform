package middleware

import (
	"mai-platform/internal/app"

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
