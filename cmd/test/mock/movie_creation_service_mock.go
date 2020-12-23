package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieCreationServiceMock struct {
	mock.Mock
}

func (mock *MovieCreationServiceMock) MovieCreation(movie model.Movie) error {
	args := mock.Called(movie)
	return args.Error(0)
}
