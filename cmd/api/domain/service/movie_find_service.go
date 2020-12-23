package service

import (
	"errors"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

type MovieFindServicePort interface {
	Find(id int64) (movie model.Movie, err error)
}

type MovieFindService struct {
	MovieRepository port.MovieRepository
}

func (movieFindService *MovieFindService) MovieFind(id int64) (movie model.Movie, err error) {

	movie, err = movieFindService.MovieRepository.Find(id)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return model.Movie{}, err
	}

	return movie, nil
}
