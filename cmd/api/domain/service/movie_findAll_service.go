package service

import (
	"errors"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

type MovieFindAllServicePort interface {
	FindAll() (movies []model.Movie, err error)
}

type MovieFindAllService struct {
	MovieRepository port.MovieRepository
}

func (movieFindAllService *MovieFindAllService) FindAll() (movies []model.Movie, err error) {

	movies, err = movieFindAllService.MovieRepository.FindAll()

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return []model.Movie{}, err
	}

	return movies, nil
}
