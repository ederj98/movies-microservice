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
	movieCreationRepository = new(mock.MovieRepositoryMock)
)

func TestWhenSendTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieCreationRepository.On("Create", movie).Times(1).Return(nil)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieCreationRepository,
	}
	err := movieCreationService.MovieCreation(movie)

	assert.Nil(t, err)
	movieCreationRepository.AssertExpectations(t)
}
func TestWhenFailedSendTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	errorExpected := errors.New("error getting repository information")
	movieCreationRepository.On("Create", movie).Times(1).Return(errorExpected)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieCreationRepository,
	}
	err := movieCreationService.MovieCreation(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieCreationRepository.AssertExpectations(t)
}
