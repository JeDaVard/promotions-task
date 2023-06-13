package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jedavard/gomotions/docs"
	"github.com/jedavard/gomotions/pkg/controllers"
	"github.com/jedavard/gomotions/pkg/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"

	_ "github.com/jedavard/gomotions/docs"
)

// @title           Promotions API
// @version         1.0
// @description     This services manages promotions

// @host      localhost:3000
// @BasePath  /api
func main() {
	setupDB()
	setupRouter()
}

func setupDB() {
	log.Println("setupDB")
	db.ConnectDatabase()
}

func setupRouter() {
	log.Println("Setup router")
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		promotions := apiV1.Group("/promotions")
		{
			promotions.POST("/bulk", controllers.UploadPromotion)
			promotions.GET("/:id", controllers.GetPromotion)

		}
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	doc := router.Group("/docs")
	{
		doc.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))
		doc.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Println("Run Router")
	router.Run()
}
