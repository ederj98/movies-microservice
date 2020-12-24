package application_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	movieUpdateServiceMock = new(mock.MovieUpdateServiceMock)
)

func TestWhenAllBeOKUpdatedMovieThenReturnNilError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieUpdateServiceMock.On("Update", movie.Id, movie).Return(nil).Once()
	movieUpdate := application.MovieUpdate{
		MovieUpdateService: movieUpdateServiceMock,
	}

	err := movieUpdate.Handler(movie.Id, movie)

	assert.Nil(t, err)
	movieUpdateServiceMock.AssertExpectations(t)
}
func TestWhenFailedUpdatedParkingThenReturnError(t *testing.T) {
	movie := builder.NewMovieDataBuilder().Build()
	expectedErrorMessage := errors.New("error getting repository information")
	movieUpdateServiceMock.On("Update", movie.Id, movie).Return(expectedErrorMessage).Once()
	movieUpdate := application.MovieUpdate{
		MovieUpdateService: movieUpdateServiceMock,
	}

	err := movieUpdate.Handler(movie.Id, movie)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieUpdateServiceMock.AssertExpectations(t)
}
