package builder

import (
	"fmt"

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

func (builder *MovieDataBuilder) BuildString() string {
	return fmt.Sprintf(
		"{\"name\":\"%s\",\"director\":\"%s\",\"writer\":\"%s\",\"stars\":\"%s\"}",
		builder.Name, builder.Director, builder.Writer, builder.Stars,
	)

}
