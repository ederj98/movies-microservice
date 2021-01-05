package service_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/stretchr/testify/assert"
)

func TestWhenDeleteTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Delete", movie.Id).Times(1).Return(nil)
	movieDeleteService := service.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.Delete(movie.Id)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}

func TestWhenFailedSendToDeleteTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Delete", movie.Id).Times(1).Return(errorExpected)
	movieDeleteService := service.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.Delete(movie.Id)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}

func TestWhenFailedFindTheMovieToDeleteThenShouldReturnError(t *testing.T) {

	movie := model.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, errorExpected)
	movieDeleteService := service.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.Delete(movie.Id)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
