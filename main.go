package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projectgo/config"
	"projectgo/handler"
	"projectgo/middleware"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the public API"})
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.GET("/users", middleware.MiddlewareToken(), userHandler.GetAll)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
