package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"projectgo/config"
	"projectgo/handler"
	"projectgo/repository"
	"projectgo/service"
)

func main() {
	database := config.InitDB()

	userRepo := repository.NewUserRepo(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.Use(gin.Recovery())

	r.POST("/register", userHandler.Register)
	r.GET("/login", userHandler.Login)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
