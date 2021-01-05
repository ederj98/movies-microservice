package service

import (
	"errors"

	"github.com/ederj98/movies-microservice/cmd/api/domain/exception"
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

const (
	errorRepository      = "error getting repository information"
	errorExistRepository = "The movie is already exist"
)

type MovieCreationServicePort interface {
	Create(movie model.Movie) (err error)
}

type MovieCreationService struct {
	MovieRepository port.MovieRepository
}

func (movieCreationService *MovieCreationService) Create(movie model.Movie) (err error) {

	exist := movieCreationService.MovieRepository.Exist(movie.Name)
	if exist != false {
		errFind := exception.DataDuplicity{ErrMessage: errorExistRepository}
		logger.Error(errorRepository, errFind)
		return errFind
	}

	err = movieCreationService.MovieRepository.Create(movie)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
