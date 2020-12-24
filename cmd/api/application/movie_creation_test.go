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
	movieCreationServiceMock = new(mock.MovieCreationServiceMock)
)

func TestWhenAllBeOKCreatingMovieThenReturnNilError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieCreationServiceMock.On("Create", movie).Return(nil).Once()
	movieCreation := application.MovieCreation{
		MovieCreationService: movieCreationServiceMock,
	}

	err := movieCreation.Handler(movie)

	assert.Nil(t, err)
	movieCreationServiceMock.AssertExpectations(t)
}
func TestWhenFailedCreatingParkingThenReturnError(t *testing.T) {
	movie := builder.NewMovieDataBuilder().Build()
	expectedErrorMessage := errors.New("error getting repository information")
	movieCreationServiceMock.On("Create", movie).Return(expectedErrorMessage).Once()
	movieCreation := application.MovieCreation{
		MovieCreationService: movieCreationServiceMock,
	}

	err := movieCreation.Handler(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieCreationServiceMock.AssertExpectations(t)
}
