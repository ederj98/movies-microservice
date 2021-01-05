package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/exception"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller/middleware"
	mockMovie "github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	movieDeleteURITest       = "/api/movies/:id"
	movieDeleteURITestWithId = "/api/movies/1"
)

var (
	movieDeleteMock mockMovie.MovieDeleteMock
)

func setupMovieDeleteController(movieDeleteMock *mockMovie.MovieDeleteMock) (*gin.Engine, *controller.MovieDeleteController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.MovieDeleteController{MovieDeleteApplication: movieDeleteMock}
}
func TestWhenMovieDeleteMovieThenReturn200(t *testing.T) {
	router, controllerMovie := setupMovieDeleteController(&movieDeleteMock)
	movieDeleteMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	movieRouterGroup := router.Group(movieDeleteURITest)
	movieRouterGroup.DELETE("", controllerMovie.MakeMovieDelete)

	response := controller.RunRequestWithHeaders(t, router, http.MethodDelete, movieDeleteURITestWithId, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	movieFindMock.AssertExpectations(t)
}

func TestWhenMovieDeleteMovieNotFoundThenReturn404(t *testing.T) {
	router, controllerMovie := setupMovieDeleteController(&movieDeleteMock)
	errorExpected := exception.DataNotFound{ErrMessage: "Movie not found"}
	movieDeleteMock.On("Handler", mock.Anything).Times(1).Return(errorExpected).Once()
	movieRouterGroup := router.Group(movieDeleteURITest)
	movieRouterGroup.DELETE("", controllerMovie.MakeMovieDelete)

	response := controller.RunRequestWithHeaders(t, router, http.MethodDelete, movieDeleteURITestWithId, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	movieFindMock.AssertExpectations(t)
}

func TestWhenMakeMovieDeleteInvalidIdThenReturn400(t *testing.T) {
	router, controllerMovie := setupMovieDeleteController(&movieDeleteMock)
	movieDeleteMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	movieRouterGroup := router.Group(movieDeleteURITest)
	movieRouterGroup.DELETE("", controllerMovie.MakeMovieDelete)

	response := controller.RunRequestWithHeaders(t, router, http.MethodDelete, movieDeleteURITest, "", nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	movieCreationMock.AssertExpectations(t)
}
