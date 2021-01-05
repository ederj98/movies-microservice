package mock

import (
	"github.com/stretchr/testify/mock"
)

type MovieDeleteRepositoryMock struct {
	mock.Mock
}

func (mock *MovieDeleteRepositoryMock) DeleteMovie(id int64) (err error) {
	args := mock.Called(id)
	return args.Error(0)
}
