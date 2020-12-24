package app

import (
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/container"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller/middleware"
	"github.com/ederj98/movies-microservice/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApplication() {

	err := godotenv.Load("../../.env")
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	router.Use(middleware.ErrorHandler())
	router.POST("/movies", container.GetMovieCreationController().MakeMovieCreation)
	router.GET("/movies", container.GetMovieFindAllController().MakeMovieFindAll)
	router.GET("/movies/:id", container.GetMovieFindController().MakeMovieFind)
	router.PUT("/movies/:id", container.GetMovieUpdateController().MakeMovieUpdate)
	router.DELETE("/movies/:id", container.GetMovieDeleteController().MakeMovieDelete)

	if err := router.Run(":8081"); err != nil {
		logger.Errorf("error running server", err)
	} else {
		logger.Info("Start the application...")
	}
}
