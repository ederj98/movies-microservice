package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
	"github.com/go-redis/redis"
)

const (
	RedisExpire = "REDIS_EXPIRE"
)

type MovieRedisRepository struct {
	Db *redis.Client
}

func (movieRedisRepository *MovieRedisRepository) Set(movie model.Movie) error {
	exp, _ := strconv.Atoi(os.Getenv(RedisExpire))
	json, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}

	if err := movieRedisRepository.Db.Set(strconv.FormatInt(movie.Id, 10), json, time.Duration(exp)*time.Second).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (movieRedisRepository *MovieRedisRepository) Get(id int64) (model.Movie, error) {

	val, err := movieRedisRepository.Db.Get(strconv.FormatInt(id, 10)).Result()
	if err != nil {
		return model.Movie{}, errors.New(fmt.Sprintf("movie with id: %d not found", id))
	}

	movie := model.Movie{}
	err = json.Unmarshal([]byte(val), &movie)
	if err != nil {
		panic(err)
	}
	return movie, nil
}
