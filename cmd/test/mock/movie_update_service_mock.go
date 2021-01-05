package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieUpdateServiceMock struct {
	mock.Mock
}

func (mock *MovieUpdateServiceMock) MovieUpdate(movie model.Movie) error {
	args := mock.Called(movie)
	return args.Error(0)
}
