package mock

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MovieFindAllServiceMock struct {
	mock.Mock
}

func (mock *MovieFindAllServiceMock) FindAllMovie() (movies []model.Movie, err error) {
	args := mock.Called()
	return args.Get(0).([]model.Movie), args.Error(1)
}
