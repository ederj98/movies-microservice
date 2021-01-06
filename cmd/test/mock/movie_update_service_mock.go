package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieUpdateServiceMock struct {
	mock.Mock
}

func (mock *MovieUpdateServiceMock) Update(id int64, movie model.Movie) error {
	args := mock.Called(id, movie)
	return args.Error(0)
}
