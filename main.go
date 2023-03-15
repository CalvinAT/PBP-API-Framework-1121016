package main

import (
	"github.com/EXPLOR/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/product", controllers.GetProducts)
	router.POST("/product", controllers.InsertProducts)
	router.PUT("/product/:prodId", controllers.UpdateProducts)
	router.DELETE("/product/:prodId", controllers.DeleteProducts)

	router.Run("localhost:6969")
}
