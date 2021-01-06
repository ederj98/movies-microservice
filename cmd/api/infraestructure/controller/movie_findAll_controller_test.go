package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller/middleware"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	mockMovie "github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	movieFindAllURITest = "/api/movies"
)

var (
	movieFindAllMock mockMovie.MovieFindAllMock
)

func setupMovieFindAllController(movieFindAllMock *mockMovie.MovieFindAllMock) (*gin.Engine, *controller.MovieFindAllController) {
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(httptest.NewRecorder())
	router.Use(middleware.ErrorHandler())
	return router, &controller.MovieFindAllController{MovieFindAllApplication: movieFindAllMock}
}
func TestWhenMovieFindAllMoviesThenReturn200AndMoviesList(t *testing.T) {
	router, controllerMovie := setupMovieFindAllController(&movieFindAllMock)
	movieList := []model.Movie{builder.NewMovieDataBuilder().Build()}
	expectedBody := "[{\"id\":1,\"name\":\"Interstellar\",\"director\":\"John Doe\",\"writer\":\"Jane Doe\",\"stars\":\"John Jr Doe, Jane M Doe\"}]"
	movieFindAllMock.On("Handler").Times(1).Return(movieList, nil).Once()
	movieFindAllRouterGroup := router.Group(movieFindAllURITest)
	movieFindAllRouterGroup.GET("", controllerMovie.MakeMovieFindAll)

	response := controller.RunRequestWithHeaders(t, router, http.MethodGet, movieFindAllURITest, "", nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, expectedBody, response.Body.String())
	movieFindAllMock.AssertExpectations(t)
}
