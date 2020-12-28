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
	movieRedisRepository repository.MovieRedisRepository
)

func TestMainRedis(m *testing.M) {
	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(containerMockServer, ctx)
	os.Exit(code)
}

func loadRedis() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := redisC.Host(ctx)
	//p, _ := redisC.MappedPort(ctx, "6379/tcp")
	_ = os.Setenv("REDIS_HOST", host)
	_ = os.Setenv("REDIS_SCHEMA", "1")

	movieRedisRepository = repository.MovieRedisRepository{
		Db: config.GetClient(),
	}
	return redisC, ctx
}

func TestMovieRedisRepository_Set(t *testing.T) {
	var movie model.Movie
	movie = builder.NewMovieDataBuilder().Build()
	err := movieRedisRepository.Set(movie)

	assert.Nil(t, err)
	assert.EqualValues(t, movie.Name, "Interstellar")
	assert.EqualValues(t, movie.Director, "John Doe")
	assert.EqualValues(t, movie.Writer, "Jane Doe")
	assert.EqualValues(t, movie.Stars, "John Jr Doe, Jane M Doe")
	assert.NotEqual(t, movie.Director, "sistemas31")
	assert.NotNil(t, movie.Id, "movie id shouldn't be nil ")
}

/*func TestMovieMysqlRepository_Find(t *testing.T) {

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
}*/
