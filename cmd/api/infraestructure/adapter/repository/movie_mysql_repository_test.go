package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/ederj98/movies-microservice/cmd/test/builder"

	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/adapter/repository"
	"github.com/ederj98/movies-microservice/cmd/api/infraestructure/config"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	movieMysqlRepository repository.MovieMySqlRepository
)

func TestMain(m *testing.M) {
	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(containerMockServer, ctx)
	os.Exit(code)
}

func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "ceiba.2020",
			"MYSQL_DATABASE":      "movies_db",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}
	mysqlC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Port()
	_ = os.Setenv("MYSQL_HOST", host)
	_ = os.Setenv("MYSQL_PORT", port)
	_ = os.Setenv("MYSQL_SCHEMA", "movies_db")
	_ = os.Setenv("MYSQL_USERNAME", "root")
	_ = os.Setenv("MYSQL_PASSWORD", "ceiba.2020")

	movieMysqlRepository = repository.MovieMySqlRepository{
		Db: config.GetDatabaseInstance(),
	}
	return mysqlC, ctx
}

func beforeAll(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
func TestMovieMysqlRepository_Create(t *testing.T) {
	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie model.Movie
	movie = builder.NewMovieDataBuilder().Build()
	err := movieMysqlRepository.Create(movie)

	assert.Nil(t, err)
	assert.EqualValues(t, movie.Name, "Interstellar", "user names are differences")
	assert.NotEqual(t, movie.Director, "sistemas31")
	assert.NotNil(t, movie.Id, "movie id shouldn't be nil ")
}

func TestUserMysqlRepository_Find(t *testing.T) {

	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie model.Movie
	movie = builder.NewMovieDataBuilder().Build()
	err := movieMysqlRepository.Create(movie)

	movieFind, err := movieMysqlRepository.Find(movie.Id)

	assert.Nil(t, err)
	assert.NotNil(t, movieFind)
	assert.EqualValues(t, movie.Id, movieFind.Id)
	assert.EqualValues(t, "John Doe", movieFind.Director)
}

func TestMovieMysqlRepository_Update(t *testing.T) {
	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie model.Movie
	movie = builder.NewMovieDataBuilder().Build()
	_ = movieMysqlRepository.Create(movie)

	movie.Director = "Jane Doe"
	errUpdate := movieMysqlRepository.Update(1, movie)

	assert.Nil(t, errUpdate)
	assert.EqualValues(t, movie.Director, "Jane Doe", "Director names are differences")
}
