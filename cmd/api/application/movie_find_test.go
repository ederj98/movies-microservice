package application_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/test/builder"
	"github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	movieFindServiceMock = new(mock.MovieFindServiceMock)
)

func TestWhenAllBeOKAndFindTheMovieThenReturnNilError(t *testing.T) {

	movieB := builder.NewMovieDataBuilder().Build()
	movieFindServiceMock.On("Find", movieB.Id).Return(movieB, nil).Once()
	movieFind := application.MovieFind{
		MovieFindService: movieFindServiceMock,
	}

	movie, err := movieFind.Handler(movieB.Id)

	assert.Nil(t, err)
	assert.Equal(t, movieB, movie)
	movieFindServiceMock.AssertExpectations(t)
}
func TestWhenFailedFindTheMovieThenReturnError(t *testing.T) {
	movieB := model.Movie{}
	expectedErrorMessage := errors.New("error getting information from service Movie List")
	movieFindServiceMock.On("Find", movieB.Id).Return(movieB, expectedErrorMessage).Once()
	movieFind := application.MovieFind{
		MovieFindService: movieFindServiceMock,
	}

	movie, err := movieFind.Handler(0)

	assert.NotNil(t, err)
	assert.Equal(t, movieB, movie)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieFindServiceMock.AssertExpectations(t)
}
