package container

import (
	"github.com/ederj98/movies-microservice/cmd/api/application"
)

func GetMovieCreationAccessApplication() application.MovieCreationApplication {
	return &application.MovieCreation{MovieCreationService: getCreateMovieService()}
}
