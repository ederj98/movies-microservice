package mapper

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/adapter/repository/entity"
)

func MovieToMovieEntity(movie model.Movie) entity.MovieEntity {
	var movieEntity entity.MovieEntity
	movieEntity.Name = movie.Name
	movieEntity.Director = movie.Director
	movieEntity.Writer = movie.Writer
	movieEntity.Stars = movie.Stars
	return movieEntity
}

func MovieEntityToMovie(movieEntity entity.MovieEntity) model.Movie {
	var movie model.Movie
	movie.Id = movieEntity.Id
	movie.Name = movieEntity.Name
	movie.Director = movieEntity.Director
	movie.Writer = movieEntity.Writer
	movie.Stars = movieEntity.Stars
	return movie
}

func MoviesEntitiesToMovies(moviesEntities []entity.MovieEntity) []model.Movie {
	var movies []model.Movie
	for _, movieEntity := range moviesEntities {
		movie := MovieEntityToMovie(movieEntity)
		movies = append(movies, movie)
	}
	return movies
}
