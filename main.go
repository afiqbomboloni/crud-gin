package main

import (
	"github.com/afiqbomboloni/crud-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/afiqbomboloni/crud-gin/controllers/productController"
)

func main(){
	r := gin.Default()

	models.ConnectionDatabase()

	r.GET("/api/v1/products", productController.Index)
	r.GET("/api/v1/products/:id", productController.Show)
	r.POST("/api/v1/products", productController.Create)
	r.PUT("/api/v1/products/:id", productController.Update)
	r.DELETE("/api/v1/products/:id", productController.Delete)

	r.Run()
}