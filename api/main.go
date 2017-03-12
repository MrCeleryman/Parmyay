package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //For Release
	router := gin.Default()
	router.Use(Cors())

	v1 := router.Group("api/v1")

	users := v1.Group("users")
	{
		users.POST("/", PostUser)
		users.GET("/", GetUsers)
		users.GET("/:id", GetUser)
		users.PUT("/:id", UpdateUser)
		users.DELETE("/:id", DeleteUser)
		users.OPTIONS("/", OptionsUser)
		users.OPTIONS("/:id", OptionsUser)
	}

	router.Run(":8900")
}
