package service_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"

	"github.com/ederj98/movies-microservice/cmd/api/domain/service"

	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/stretchr/testify/assert"
)

func TestWhenGetTheMovieListFromRepositoryThenShouldReturnOk(t *testing.T) {
	movieLots := []model.Movie{builder.NewMovieDataBuilder().Build()}
	movieRepository.On("FindAll").Times(1).Return(movieLots, nil)
	movieFindAllService := service.MovieFindAllService{
		MovieRepository: movieRepository,
	}

	movies, err := movieFindAllService.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, movieLots, movies)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedGetTheMovieListFromRepositoryThenShouldReturnError(t *testing.T) {

	movieLots := []model.Movie{}
	errorExpected := errors.New(errorRepository)
	movieRepository.On("FindAll").Times(1).Return(movieLots, errorExpected)
	movieFindAllService := service.MovieFindAllService{
		MovieRepository: movieRepository,
	}

	movies, err := movieFindAllService.FindAll()

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	assert.Equal(t, movieLots, movies)
	movieRepository.AssertExpectations(t)
}
