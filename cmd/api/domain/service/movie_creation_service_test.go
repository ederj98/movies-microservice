package service_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/service"

	"github.com/ederj98/movies-microservice/cmd/test/builder"

	"github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	movieRepository = new(mock.MovieRepositoryMock)
)

const (
	errorRepository = "error getting repository information"
)

func TestWhenSendTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieRepository.On("Create", movie).Times(1).Return(nil)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.MovieCreation(movie)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedSendTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Create", movie).Times(1).Return(errorExpected)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.MovieCreation(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
