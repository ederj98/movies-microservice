package mock

import "github.com/stretchr/testify/mock"

type MovieDeleteServiceMock struct {
	mock.Mock
}

func (mock *MovieDeleteServiceMock) MovieDelete(id int64) error {
	args := mock.Called(id)
	return args.Error(0)
}
