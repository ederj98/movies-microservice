package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieFindRepositoryMock struct {
	mock.Mock
}

func (mock *MovieFindRepositoryMock) FindMovie(id int64) (movie model.Movie, err error) {
	args := mock.Called(id)
	return args.Get(0).(model.Movie), args.Error(1)
}
