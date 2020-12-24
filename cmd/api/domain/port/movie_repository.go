package port

import "github.com/ederj98/movies-microservice/cmd/api/domain/model"

type MovieRepository interface {
	Create(model.Movie) error
	Exist(string) bool
	Find(int64) (model.Movie, error)
	FindAll() ([]model.Movie, error)
	Update(int64, model.Movie) error
	Delete(int64) error
}
