package service_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/stretchr/testify/assert"
)

const (
	errorFindRepository = "Movie not found"
)

func TestWhenGetTheMovieFindFromRepositoryThenShouldReturnOk(t *testing.T) {
	movieB := builder.NewMovieDataBuilder().Build()
	movieR := model.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieRedisRepository.On("Get", movieB.Id).Times(1).Return(movieR, errorExpected)
	movieRepository.On("Find", movieB.Id).Times(1).Return(movieB, nil)
	movieRedisRepository.On("Set", movieB).Times(1).Return(nil)
	movieFindService := service.MovieFindService{
		MovieRepository:      movieRepository,
		MovieRedisRepository: movieRedisRepository,
	}

	movie, err := movieFindService.Find(movieB.Id)

	assert.Nil(t, err)
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedGetTheMovieFromRepositoryThenShouldReturnError(t *testing.T) {

	movieB := model.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieR := model.Movie{}
	movieRedisRepository.On("Get", movieB.Id).Times(1).Return(movieR, errorExpected)
	movieRepository.On("Find", movieB.Id).Times(1).Return(movieB, errorExpected)
	movieFindService := service.MovieFindService{
		MovieRepository:      movieRepository,
		MovieRedisRepository: movieRedisRepository,
	}

	movie, err := movieFindService.Find(0)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}

func TestWhenGetTheMovieFindFromRedisRepositoryThenShouldReturnOk(t *testing.T) {
	movieB := builder.NewMovieDataBuilder().Build()
	movieRedisRepository.On("Get", movieB.Id).Times(1).Return(movieB, nil)
	movieFindService := service.MovieFindService{
		MovieRepository:      movieRepository,
		MovieRedisRepository: movieRedisRepository,
	}

	movie, err := movieFindService.Find(movieB.Id)

	assert.Nil(t, err)
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}
