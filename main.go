package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"projectgo/config"
	"projectgo/handler"
	"projectgo/repository"
	"projectgo/service"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{w, http.StatusCreated}
		next.ServeHTTP(lrw, r)

		log.Printf(" %s %d %s", r.Method, lrw.statusCode, time.Since(start))
	})
}

func main() {
	database := config.InitDB()

	userRepo := repository.NewUserRepo(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.Use(gin.Recovery())

	r.POST("/register", userHandler.Register)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
