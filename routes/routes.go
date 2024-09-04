package routes

import (
	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	return router
}
