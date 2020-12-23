package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieUpdateRepositoryMock struct {
	mock.Mock
}

func (mock *MovieUpdateRepositoryMock) UpdateMovie(movie model.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}
