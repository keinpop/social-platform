package main

import (
	"log"
	"mai-platform/api"
	"mai-platform/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

const CONFIG_PATH = "./etc/config.yaml"

func main() {
	c, err := api.NewConfig(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	a := api.NewApp(c)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/api/technology/list", a.GetTechnologies)
	r.POST("/api/technology", a.AddTechnology)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
