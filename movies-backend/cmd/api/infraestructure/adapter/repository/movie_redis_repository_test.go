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
	p, _ := redisC.MappedPort(ctx, "6379/tcp")
	port := p.Port()
	_ = os.Setenv("REDIS_HOST", host)
	_ = os.Setenv("REDIS_PORT", port)
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
}

func TestMovieRedisRepository_Get(t *testing.T) {

	var movie model.Movie
	movie = builder.NewMovieDataBuilder().Build()
	err := movieRedisRepository.Set(movie)

	movieFind, err := movieRedisRepository.Get(movie.Id)

	assert.Nil(t, err)
	assert.NotNil(t, movieFind)
	assert.EqualValues(t, movie.Id, movieFind.Id)
	assert.EqualValues(t, movie.Name, movieFind.Name)
	assert.EqualValues(t, movie.Director, movieFind.Director)
	assert.EqualValues(t, movie.Writer, movieFind.Writer)
	assert.EqualValues(t, movie.Stars, movieFind.Stars)
}
