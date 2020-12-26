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
	errorRepository      = "error getting repository information"
	errorExistRepository = "The movie is already exist"
)

func TestWhenSendTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieRepository.On("Exist", movie.Name).Times(1).Return(false)
	movieRepository.On("Create", movie).Times(1).Return(nil)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.Create(movie)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedSendTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Exist", movie.Name).Times(1).Return(false)
	movieRepository.On("Create", movie).Times(1).Return(errorExpected)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.Create(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}

func TestWhenMovieExistThenShouldReturnError(t *testing.T) {
	movie := builder.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorExistRepository)
	movieRepository.On("Exist", movie.Name).Times(1).Return(true)
	movieCreationService := service.MovieCreationService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.Create(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
