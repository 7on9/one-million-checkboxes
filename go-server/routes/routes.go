package routes

import (
	"one-million-checkboxes/services"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	checkboxesController "one-million-checkboxes/controllers/checkboxes"
	docs "one-million-checkboxes/docs"

	swaggerfiles "github.com/swaggo/files"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	services.InitServices()
	// r.GET("/hello/:name", getWithParams)

	checkboxes := r.Group("/api/v1/checkboxes")
	{
		checkboxes.POST("/update", checkboxesController.UpdateBitSet)
		checkboxes.GET("/current", checkboxesController.GetCurrentBitSet)
	}

	// Swagger
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Title = "One Million Checkboxes"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "This is a simple API to manage a bitset of one million checkboxes"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
