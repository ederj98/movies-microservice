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
	movieFindAllRepositoryMock = new(mock.MovieRepositoryMock)
)

func TestWhenAllBeOKAndListMovieThenReturnNilError(t *testing.T) {

	movieLots := []model.Movie{builder.NewMovieDataBuilder().Build()}
	movieFindAllRepositoryMock.On("FindAll").Return(movieLots, nil).Once()
	movieFindAll := application.MovieFindAll{
		MovieRepository: movieFindAllRepositoryMock,
	}

	parking, err := movieFindAll.Handler()

	assert.Nil(t, err)
	assert.Equal(t, movieLots, parking)
	movieFindAllRepositoryMock.AssertExpectations(t)
}
func TestWhenFailedListMovieThenReturnError(t *testing.T) {
	movieLots := []model.Movie{}
	expectedErrorMessage := errors.New("error getting information from service Movie List")
	movieFindAllRepositoryMock.On("FindAll").Return(movieLots, expectedErrorMessage).Once()
	movieFindAll := application.MovieFindAll{
		MovieRepository: movieFindAllRepositoryMock,
	}

	movies, err := movieFindAll.Handler()

	assert.NotNil(t, err)
	assert.Equal(t, movieLots, movies)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieFindAllRepositoryMock.AssertExpectations(t)
}
