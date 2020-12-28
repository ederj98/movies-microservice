package port

import "github.com/ederj98/movies-microservice/cmd/api/domain/model"

type MovieRedisRepository interface {
	Set(model.Movie) error
	Get(int64) (model.Movie, error)
}
