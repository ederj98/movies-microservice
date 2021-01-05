package controller

import (
	"net/http"
	"strconv"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/marshaller"
	"github.com/ederj98/movies-microservice/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

const (
	failedFindMovie = "we didn't the movie"
)

// FindMovieController used for inject the use case
type MovieFindController struct {
	MovieFindApplication application.MovieFindApplication
}

//MakeMovieFind is to execute the use case for get a movie
func (movieFindController *MovieFindController) MakeMovieFind(context *gin.Context) {
	id := movieFindController.mapMovieFind(context)
	movie, err := movieFindController.MovieFindApplication.Handler(id)

	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, marshaller.Marshall(movie))

}

// mapMovieFind used for get attributes for  any movie from url
func (movieFindController *MovieFindController) mapMovieFind(c *gin.Context) (id int64) {
	id, movieErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if movieErr != nil {
		idErr := apierrors.NewApiError("Invalid id", movieErr.Error(), 400, nil)
		c.JSON(idErr.Status(), idErr)
		return
	}
	return id
}
