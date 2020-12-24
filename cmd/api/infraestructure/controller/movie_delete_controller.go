package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

const (
	MovieDeletedMsg = "the Movie with id: %v was deleted successfully"
)

// MovieDeleteController used for inject the use case
type MovieDeleteController struct {
	MovieDeleteApplication application.MovieDeleteApplication
}

//MakeMovieDelete is to execute the use case for update the movie
func (movieDeleteController *MovieDeleteController) MakeMovieDelete(context *gin.Context) {

	id := movieDeleteController.mapMovieDelete(context)
	err := movieDeleteController.MovieDeleteApplication.Handler(id)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, fmt.Sprintf(MovieDeletedMsg, id))

}

// mapMovie used for get attributes for  any movie from url
func (movieDeleteController *MovieDeleteController) mapMovieDelete(c *gin.Context) (id int64) {
	id, movieErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if movieErr != nil {
		idErr := apierrors.NewApiError("Invalid id", movieErr.Error(), 400, nil)
		c.JSON(idErr.Status(), idErr)
		return
	}
	return id
}
