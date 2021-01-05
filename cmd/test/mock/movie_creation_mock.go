package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieCreationMock struct {
	mock.Mock
}

func (mock *MovieCreationMock) Handler(movie model.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}
