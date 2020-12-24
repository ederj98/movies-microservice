package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

const (
	MovieUpdatedMsg = "the Movie with id: %v was updated successfully"
)

// MovieUpdateController used for inject the use case
type MovieUpdateController struct {
	MovieUpdateApplication application.MovieUpdateApplication
}

//MakeMovieUpdate is to execute the use case for update the movie
func (movieUpdateController *MovieUpdateController) MakeMovieUpdate(context *gin.Context) {

	id, movie := movieUpdateController.mapMovieUpdate(context)
	err := movieUpdateController.MovieUpdateApplication.Handler(id, movie)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusNoContent, fmt.Sprintf(MovieUpdatedMsg, id))

}

// mapMovie used for get attributes for  any movie from url
func (movieUpdateController *MovieUpdateController) mapMovieUpdate(c *gin.Context) (id int64, movie model.Movie) {
	id, movieErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if movieErr != nil {
		idErr := apierrors.NewApiError("Invalid id", movieErr.Error(), 400, nil)
		c.JSON(idErr.Status(), idErr)
		return
	}

	var movieL model.Movie
	if err := c.ShouldBindJSON(&movieL); err != nil {
		restErr := apierrors.NewApiError("Invalid json", err.Error(), 400, nil)
		c.JSON(restErr.Status(), restErr)
		return
	}
	return id, movieL
}
