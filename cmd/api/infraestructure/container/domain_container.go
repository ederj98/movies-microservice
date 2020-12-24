package container

import "github.com/ederj98/movies-microservice/cmd/api/domain/service"

func getCreateMovieService() service.MovieCreationServicePort {
	return &service.MovieCreationService{
		MovieRepository: getMovieRepository(),
	}
}

func getUpdateMovieService() service.MovieUpdateServicePort {
	return &service.MovieUpdateService{
		MovieRepository: getMovieRepository(),
	}
}

func getDeleteMovieService() service.MovieDeleteServicePort {
	return &service.MovieDeleteService{
		MovieRepository: getMovieRepository(),
	}
}

func getFindAllMovieService() service.MovieFindAllServicePort {
	return &service.MovieFindAllService{
		MovieRepository: getMovieRepository(),
	}
}

func getFindMovieService() service.MovieFindServicePort {
	return &service.MovieFindService{
		MovieRepository: getMovieRepository(),
	}
}
