package main

import (
	"go-jwt/pkg/controllers"
	"go-jwt/pkg/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)

	r.Run()
}
