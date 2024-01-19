package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xiayulin123/GO_Ecommerce/controllers"
	"github.com/xiayulin123/GO_Ecommerce/database"
	"github.com/xiayulin123/GO_Ecommerce/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.new()

	router.Use(gin.Logger())

	routers.UserRoutes(router)

	router.Use(middleware.Authentication())
}
