package builder

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
)

type MovieDataBuilder struct {
	Id       int64
	Name     string
	Director string
	Writer   string
	Stars    string
}

func NewMovieDataBuilder() *MovieDataBuilder {
	return &MovieDataBuilder{
		Id:       1,
		Name:     "Interstellar",
		Director: "John Doe",
		Writer:   "Jane Doe",
		Stars:    "John Jr Doe, Jane M Doe",
	}
}
func (builder *MovieDataBuilder) Build() model.Movie {
	return model.Movie{
		Id:       builder.Id,
		Name:     builder.Name,
		Director: builder.Director,
		Writer:   builder.Writer,
		Stars:    builder.Stars,
	}
}
