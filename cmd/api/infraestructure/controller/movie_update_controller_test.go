package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/exception"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller/middleware"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	mockMovie "github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	movieUpdateURITest       = "/api/movies/:id"
	movieUpdateURITestWithId = "/api/movies/1"
)

var (
	movieUpdateMock mockMovie.MovieUpdateMock
)

func setupMovieUpdateController(movieUpdateMock *mockMovie.MovieUpdateMock) (*gin.Engine, *controller.MovieUpdateController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.MovieUpdateController{MovieUpdateApplication: movieUpdateMock}
}
func TestWhenMakeMovieUpdateThenReturn204(t *testing.T) {
	router, controllerMovie := setupMovieUpdateController(&movieUpdateMock)
	movie := builder.NewMovieDataBuilder().BuildString()
	movieUpdateMock.On("Handler", mock.Anything, mock.Anything).Times(1).Return(nil).Once()
	movieRouterGroup := router.Group(movieUpdateURITest)
	movieRouterGroup.PUT("", controllerMovie.MakeMovieUpdate)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPut, movieUpdateURITestWithId, movie, nil)

	assert.Equal(t, http.StatusNoContent, response.Code)
	movieUpdateMock.AssertExpectations(t)
}

func TestWhenMakeMovieUpdateMovieNotFoundThenReturn404(t *testing.T) {
	router, controllerMovie := setupMovieUpdateController(&movieUpdateMock)
	errorExpected := exception.DataNotFound{ErrMessage: "Movie not found"}
	movie := builder.NewMovieDataBuilder().BuildString()
	movieUpdateMock.On("Handler", mock.Anything, mock.Anything).Times(1).Return(errorExpected).Once()
	movieRouterGroup := router.Group(movieUpdateURITest)
	movieRouterGroup.PUT("", controllerMovie.MakeMovieUpdate)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPut, movieUpdateURITestWithId, movie, nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	movieCreationMock.AssertExpectations(t)
}

func TestWhenMakeMovieUpdateInvalidIdThenReturn400(t *testing.T) {
	router, controllerMovie := setupMovieUpdateController(&movieUpdateMock)
	movie := builder.NewMovieDataBuilder().BuildString()
	movieUpdateMock.On("Handler", mock.Anything, mock.Anything).Times(1).Return(nil).Once()
	movieRouterGroup := router.Group(movieUpdateURITest)
	movieRouterGroup.PUT("", controllerMovie.MakeMovieUpdate)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPut, movieUpdateURITest, movie, nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	movieCreationMock.AssertExpectations(t)
}
