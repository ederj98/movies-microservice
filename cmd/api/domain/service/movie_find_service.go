package service

import (
	"fmt"

	"github.com/ederj98/movies-microservice/cmd/api/domain/exception"
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/pkg/logger"
)

const (
	errorNotFoundRepository = "The movie isn't not found"
)

type MovieFindServicePort interface {
	Find(id int64) (movie model.Movie, err error)
}

type MovieFindService struct {
	MovieRepository      port.MovieRepository
	MovieRedisRepository port.MovieRedisRepository
}

func (movieFindService *MovieFindService) Find(id int64) (movie model.Movie, err error) {

	movie, err = movieFindService.MovieRedisRepository.Get(id)
	logger.Info(fmt.Sprintf("Consulta Redis: %s", movie))

	if err != nil {
		movie, err = movieFindService.MovieRepository.Find(id)
		if err != nil {
			err = exception.DataNotFound{ErrMessage: errorNotFoundRepository}
			logger.Error(errorRepository, err)
			return model.Movie{}, err
		}
		error := movieFindService.MovieRedisRepository.Set(movie)
		if error != nil {
			logger.Error(errorRepository, error)
		}
		return movie, nil
	}

	return movie, nil
}
