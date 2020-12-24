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

func getMovieRepository() port.MovieRepository {
	return &repository.MovieMySqlRepository{
		Db: config.GetDatabaseInstance(),
	}
}
