package service

import (
	"errors"

	"github.com/ederj98/movies-microservice/cmd/api/domain/exception"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

type MovieDeleteServicePort interface {
	Delete(id int64) (err error)
}

type MovieDeleteService struct {
	MovieRepository port.MovieRepository
}

func (movieDeleteService *MovieDeleteService) Delete(id int64) (err error) {

	_, exist := movieDeleteService.MovieRepository.Find(id)
	if exist != nil {
		err = exception.DataNotFound{ErrMessage: errorNotFoundRepository}
		logger.Error(errorRepository, err)
		return err
	}

	err = movieDeleteService.MovieRepository.Delete(id)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
