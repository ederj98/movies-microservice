package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/exception"
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller/middleware"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	mockMovie "github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	movieFindURITest       = "/api/movies/:id"
	movieFindURITestWithId = "/api/movies/1"
)

var (
	movieFindMock mockMovie.MovieFindMock
)

func setupMovieFindController(movieFindMock *mockMovie.MovieFindMock) (*gin.Engine, *controller.MovieFindController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.MovieFindController{MovieFindApplication: movieFindMock}
}
func TestWhenMovieFindMovieThenReturn200AndMovie(t *testing.T) {
	router, controllerMovie := setupMovieFindController(&movieFindMock)
	movie := builder.NewMovieDataBuilder().Build()
	movieFindMock.On("Handler", mock.Anything).Times(1).Return(movie, nil).Once()
	movieFindRouterGroup := router.Group(movieFindURITest)
	movieFindRouterGroup.GET("", controllerMovie.MakeMovieFind)
	fmt.Println(movieFindRouterGroup.BasePath())

	response := controller.RunRequestWithHeaders(t, router, http.MethodGet, movieFindURITestWithId, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	movieFindMock.AssertExpectations(t)
}

func TestWhenMovieFindMovieThenReturn404(t *testing.T) {
	router, controllerMovie := setupMovieFindController(&movieFindMock)
	movie := model.Movie{}
	errorExpected := exception.DataNotFound{ErrMessage: "Movie not found"}
	movieFindMock.On("Handler", mock.Anything).Times(1).Return(movie, errorExpected).Once()
	movieFindRouterGroup := router.Group(movieFindURITest)
	movieFindRouterGroup.GET("", controllerMovie.MakeMovieFind)

	response := controller.RunRequestWithHeaders(t, router, http.MethodGet, movieFindURITestWithId, "", nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
	movieFindMock.AssertExpectations(t)
}

func TestWhenMakeMovieFindInvalidIdThenReturn400(t *testing.T) {
	router, controllerMovie := setupMovieFindController(&movieFindMock)
	movie := model.Movie{}
	movieFindMock.On("Handler", mock.Anything).Times(1).Return(movie, nil).Once()
	movieRouterGroup := router.Group(movieFindURITest)
	movieRouterGroup.GET("", controllerMovie.MakeMovieFind)

	response := controller.RunRequestWithHeaders(t, router, http.MethodGet, movieFindURITest, "", nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	movieCreationMock.AssertExpectations(t)
}
