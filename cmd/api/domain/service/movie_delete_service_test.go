package service_test

import (
	"testing"

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
