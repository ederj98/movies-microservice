package controller

import (
	"net/http"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/util"
	"github.com/ederj98/movies-microservice/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

const (
	failedFindAllMovies = "we didn't list all movies"
)

// FinAllMoviesController used for inject the use case
type MovieFindAllController struct {
	MovieFindAllApplication application.MovieFindAllApplication
}

//MakeMovieFindAll is to execute the use case for get all movies
func (movieFindAllController *MovieFindAllController) MakeMovieFindAll(context *gin.Context) {

	moviesList, err := movieFindAllController.MovieFindAllApplication.Handler()

	if err != nil {
		err := apierrors.NewNotFoundApiError(failedFindAllMovies)
		context.JSON(err.Status(), err)
	}

	//Resty usage test
	util.RestyGet()

	context.JSON(http.StatusOK, moviesList)
}
