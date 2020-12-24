package service_test

import (
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/service"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/stretchr/testify/assert"
)

func TestWhenDeleteTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builder.NewMovieDataBuilder().Build()
	movieRepository.On("Delete", movie.Id).Times(1).Return(nil)
	movieDeleteService := service.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.MovieDelete(movie.Id)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}
