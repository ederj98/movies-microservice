package controller

import (
	"fmt"
	"net/http"

	"github.com/ederj98/movies-microservice/pkg/apierrors"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/gin-gonic/gin"
)

const (
	MovieCreatedMsg = "the Movie %v was created successfully"
)

// MovieCreationController  used for inject the use case
type MovieCreationController struct {
	MovieCreationApplication application.MovieCreationApplication
}

//MakeMovieCreation is to execute the use case for create the movie
func (movieCreationController *MovieCreationController) MakeMovieCreation(context *gin.Context) {

	movie := movieCreationController.mapMovie(context)
	err := movieCreationController.MovieCreationApplication.Handler(movie)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusCreated, fmt.Sprintf(MovieCreatedMsg, movie.Name))

}

// mapMovie used for get attributes for  any movie from url
func (movieCreationController *MovieCreationController) mapMovie(c *gin.Context) (movie model.Movie) {

	var movieL model.Movie
	if err := c.ShouldBindJSON(&movieL); err != nil {
		restErr := apierrors.NewApiError("invalid json", err.Error(), 400, nil)
		c.JSON(restErr.Status(), restErr)
		return
	}
	return movieL
}

func abort(ctx *gin.Context, err error) {
	ctx.Error(err)
	ctx.Abort()
}
