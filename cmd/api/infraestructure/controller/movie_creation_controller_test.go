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
	movieCreationURITest = "/api/movies"
)

var (
	movieCreationMock mockMovie.MovieCreationMock
)

func setupMovieCreationController(movieCreationMock *mockMovie.MovieCreationMock) (*gin.Engine, *controller.MovieCreationController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.MovieCreationController{MovieCreationApplication: movieCreationMock}
}
func TestWhenMakeMovieCreationThenReturn201(t *testing.T) {
	movie := builder.NewMovieDataBuilder().BuildString()
	router, controllerMovie := setupMovieCreationController(&movieCreationMock)
	movieCreationMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	movieRouterGroup := router.Group(movieCreationURITest)
	movieRouterGroup.POST("", controllerMovie.MakeMovieCreation)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, movieCreationURITest, movie, nil)

	assert.Equal(t, http.StatusCreated, response.Code)
	movieCreationMock.AssertExpectations(t)
}

func TestWhenMakeMovieCreationFoundThenReturn404(t *testing.T) {
	router, controllerMovie := setupMovieCreationController(&movieCreationMock)
	errorExpected := exception.DataNotFound{ErrMessage: "The movie is already exist"}
	movie := builder.NewMovieDataBuilder().BuildString()
	movieCreationMock.On("Handler", mock.Anything).Times(1).Return(errorExpected).Once()
	movieRouterGroup := router.Group(movieCreationURITest)
	movieRouterGroup.POST("", controllerMovie.MakeMovieCreation)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, movieCreationURITest, movie, nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	movieCreationMock.AssertExpectations(t)
}
