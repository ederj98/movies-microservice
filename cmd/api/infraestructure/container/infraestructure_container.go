package container

import (
	"github.com/ederj98/movies-microservice/cmd/api/domain/port"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/adapter/repository"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/config"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/controller"
)

func GetMovieCreationController() *controller.MovieCreationController {
	return &controller.MovieCreationController{MovieCreationApplication: GetMovieCreationAccessApplication()}
}

func GetMovieUpdateController() *controller.MovieUpdateController {
	return &controller.MovieUpdateController{MovieUpdateApplication: GetMovieUpdateAccessApplication()}
}

func GetMovieDeleteController() *controller.MovieDeleteController {
	return &controller.MovieDeleteController{MovieDeleteApplication: GetMovieDeleteAccessApplication()}
}

func GetMovieFindAllController() *controller.MovieFindAllController {
	return &controller.MovieFindAllController{MovieFindAllApplication: GetMovieFindAllAccessApplication()}
}

func GetMovieFindController() *controller.MovieFindController {
	return &controller.MovieFindController{MovieFindApplication: GetMovieFindAccessApplication()}
}

func getMovieRepository() port.MovieRepository {
	return &repository.MovieMySqlRepository{
		Db: config.GetDatabaseInstance(),
	}
}
