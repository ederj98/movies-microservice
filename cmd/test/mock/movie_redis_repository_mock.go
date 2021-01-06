package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieRedisRepositoryMock struct {
	mock.Mock
}

func (mock *MovieRedisRepositoryMock) Set(movie model.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}

func (mock *MovieRedisRepositoryMock) Get(id int64) (movie model.Movie, err error) {
	args := mock.Called(id)
	return args.Get(0).(model.Movie), args.Error(1)
}
