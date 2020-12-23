package service

import (
	"errors"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

type MovieUpdateServicePort interface {
	Update(id int64, movie model.Movie) (err error)
}

type MovieUpdateService struct {
	MovieRepository port.MovieRepository
}

func (movieUpdateService *MovieUpdateService) MovieUpdate(id int64, movie model.Movie) (err error) {

	err = movieUpdateService.MovieRepository.Update(id, movie)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
