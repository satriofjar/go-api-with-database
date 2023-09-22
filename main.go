package main

import (
	productcontroller "go-api-with-database/controllers/productController"
	"go-api-with-database/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDatabase()

	route.GET("/api/products", productcontroller.Index)
	route.GET("/api/product/:id", productcontroller.Show)
	route.POST("/api/product", productcontroller.Create)
	route.PUT("/api/product/:id", productcontroller.Update)
	route.DELETE("/api/product", productcontroller.Delete)

	route.Run(":8000")
}
