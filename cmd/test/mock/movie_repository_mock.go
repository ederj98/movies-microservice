package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieRepositoryMock struct {
	mock.Mock
}

func (mock *MovieRepositoryMock) Create(movie model.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}

func (mock *MovieRepositoryMock) Find(id int64) (movie model.Movie, err error) {
	args := mock.Called(id)
	return args.Get(0).(model.Movie), args.Error(1)
}

func (mock *MovieRepositoryMock) FindAll() (movies []model.Movie, err error) {
	args := mock.Called()
	return args.Get(0).([]model.Movie), args.Error(1)
}

func (mock *MovieRepositoryMock) Update(id int64, movie model.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}

func (mock *MovieRepositoryMock) Delete(id int64) (err error) {
	args := mock.Called(id)
	return args.Error(0)
}
