package container

import (
	"github.com/ederj98/movies-microservice/cmd/api/application"
)

func GetMovieCreationAccessApplication() application.MovieCreationApplication {
	return &application.MovieCreation{MovieCreationService: getCreateMovieService()}
}

func GetMovieUpdateAccessApplication() application.MovieUpdateApplication {
	return &application.MovieUpdate{MovieUpdateService: getUpdateMovieService()}
}

func GetMovieDeleteAccessApplication() application.MovieDeleteApplication {
	return &application.MovieDelete{MovieDeleteService: getDeleteMovieService()}
}

func GetMovieFindAllAccessApplication() application.MovieFindAllApplication {
	return &application.MovieFindAll{MovieFindAllService: getFindAllMovieService()}
}

func GetMovieFindAccessApplication() application.MovieFindApplication {
	return &application.MovieFind{MovieFindService: getFindMovieService()}
}
