package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller/middleware"
	mockMovie "github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	movieCreationURITest = "/api"
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
	//movie := builder.NewMovieDataBuilder().Build()
	router, controllerMovie := setupMovieCreationController(&movieCreationMock)
	movieCreationMock.On("Handler", mock.Anything).Times(1).Return(nil).Once()
	movieRouterGroup := router.Group(movieCreationURITest)
	movieRouterGroup.POST("/movies", controllerMovie.MakeMovieCreation)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, movieCreationURITest, "", nil)
	fmt.Println(response)
	assert.Equal(t, http.StatusCreated, response.Code)
	movieCreationMock.AssertExpectations(t)
}

/*func TestWhenMakeParkingCreationFailedThenReturn400(t *testing.T) {

	router, controllerMovie := setupMovieCreationController(&movieCreationMock)
	errorExpected := exception.DataNotFound{ErrMessage: "we didn't found information"}
	movieCreationMock.On("Handler", mock.Anything).Times(1).Return(errorExpected).Once()
	movieRouterGroup := router.Group(movieCreationURITest)
	movieRouterGroup.POST("", controllerMovie.MakeMovieCreation)

	response := controller.RunRequestWithHeaders(t, router, http.MethodPost, movieCreationURITest, "", nil)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	movieCreationMock.AssertExpectations(t)
}*/
