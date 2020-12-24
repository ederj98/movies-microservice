package service_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/stretchr/testify/assert"
)

func TestWhenSendToUpdateTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Update", movie.Id, movie).Times(1).Return(nil)
	movieUpdateService := service.MovieUpdateService{
		MovieRepository: movieRepository,
	}
	err := movieUpdateService.Update(movie.Id, movie)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedSendToUpdateTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Update", movie.Id, movie).Times(1).Return(errorExpected)
	movieUpdateService := service.MovieUpdateService{
		MovieRepository: movieRepository,
	}
	err := movieUpdateService.Update(movie.Id, movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
