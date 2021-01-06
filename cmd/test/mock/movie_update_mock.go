package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieUpdateMock struct {
	mock.Mock
}

func (mock *MovieUpdateMock) Handler(id int64, movie model.Movie) (err error) {
	args := mock.Called(id, movie)
	return args.Error(0)
}
