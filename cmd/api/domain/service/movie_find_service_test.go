package service_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/stretchr/testify/assert"
)

func TestWhenGetTheMovieFindFromRepositoryThenShouldReturnOk(t *testing.T) {
	movieB := builder.NewMovieDataBuilder().Build()
	movieRepository.On("Find", movieB.Id).Times(1).Return(movieB, nil)
	movieFindService := service.MovieFindService{
		MovieRepository: movieRepository,
	}

	movie, err := movieFindService.MovieFind(movieB.Id)

	assert.Nil(t, err)
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedGetTheMovieFromRepositoryThenShouldReturnError(t *testing.T) {

	movieB := model.Movie{}
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Find", movieB.Id).Times(1).Return(movieB, errorExpected)
	movieFindService := service.MovieFindService{
		MovieRepository: movieRepository,
	}

	movie, err := movieFindService.MovieFind(0)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}
