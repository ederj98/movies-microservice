package application_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	movieFindAllServiceMock = new(mock.MovieFindAllServiceMock)
)

func TestWhenAllBeOKAndListMovieThenReturnNilError(t *testing.T) {

	movieLots := []model.Movie{builder.NewMovieDataBuilder().Build()}
	movieFindAllServiceMock.On("FindAll").Return(movieLots, nil).Once()
	movieFindAll := application.MovieFindAll{
		MovieFindAllService: movieFindAllServiceMock,
	}

	parking, err := movieFindAll.Handler()

	assert.Nil(t, err)
	assert.Equal(t, movieLots, parking)
	movieFindAllServiceMock.AssertExpectations(t)
}
func TestWhenFailedListMovieThenReturnError(t *testing.T) {
	movieLots := []model.Movie{}
	expectedErrorMessage := errors.New("error getting information from service Movie List")
	movieFindAllServiceMock.On("FindAll").Return(movieLots, expectedErrorMessage).Once()
	movieFindAll := application.MovieFindAll{
		MovieFindAllService: movieFindAllServiceMock,
	}

	movies, err := movieFindAll.Handler()

	assert.NotNil(t, err)
	assert.Equal(t, movieLots, movies)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieFindAllServiceMock.AssertExpectations(t)
}
