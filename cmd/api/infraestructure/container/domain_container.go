package container

import "github.com/ederj98/movies-microservice/cmd/api/domain/service"

func getCreateMovieService() service.MovieCreationServicePort {
	return &service.MovieCreationService{
		MovieRepository: getMovieRepository(),
	}
}
