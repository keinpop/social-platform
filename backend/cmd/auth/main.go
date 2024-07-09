package main

import (
	"fmt"
	"log"
	"net/http"

	"mai-platform/internal/auth"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const CONFIG_PATH = "./config/auth_config.yaml"

func main() {
	c, err := auth.NewConfig(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	a := auth.NewAuth(c)
	if err := a.Init(); err != nil {
		log.Fatalf("Failed to init auth: %v", err)
	}

	engine := gin.Default()
	// the jwt middleware
	authMiddleware, err := jwt.New(a.GetJWTMiddleware())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// register middleware
	engine.Use(handlerMiddleWare(authMiddleware))

	// register route
	registerRoute(engine, authMiddleware, a)

	// start http server
	if err = http.ListenAndServe(":"+fmt.Sprint(c.Port), engine); err != nil {
		log.Fatal(err)
	}
}

func registerRoute(r *gin.Engine, handle *jwt.GinJWTMiddleware, a *auth.Auth) {
	r.POST("/login", handle.LoginHandler)
	r.POST("/register", a.Register)

	auth := r.Group("/auth", handle.MiddlewareFunc())
	auth.GET("/refresh_token", handle.RefreshHandler)
	auth.GET("/check", a.CheckToken)
}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}
