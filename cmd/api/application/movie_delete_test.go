package application_test

import (
	"errors"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/application"
	"github.com/ederj98/movies-microservice/cmd/test/mock"
	"github.com/stretchr/testify/assert"
)

var (
	movieDeleteServiceMock = new(mock.MovieDeleteServiceMock)
)

func TestWhenAllBeOKDeletedMovieThenReturnNilError(t *testing.T) {
	const id int64 = 5
	movieDeleteServiceMock.On("Delete", id).Return(nil).Once()
	movieDelete := application.MovieDelete{
		MovieDeleteService: movieDeleteServiceMock,
	}

	err := movieDelete.Handler(id)

	assert.Nil(t, err)
	movieDeleteServiceMock.AssertExpectations(t)
}
func TestWhenFailedDeletedParkingThenReturnError(t *testing.T) {
	const id int64 = 5
	expectedErrorMessage := errors.New("error getting repository information")
	movieDeleteServiceMock.On("Delete", id).Return(expectedErrorMessage).Once()
	movieDelete := application.MovieDelete{
		MovieDeleteService: movieDeleteServiceMock,
	}

	err := movieDelete.Handler(id)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieDeleteServiceMock.AssertExpectations(t)
}
