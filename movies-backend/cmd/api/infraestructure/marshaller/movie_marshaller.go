package marshaller

import (
	"encoding/json"
	"fmt"

	"github.com/ederj98/movies-microservice/cmd/api/domain/model"
)

type MovieJson struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Stars    string `json:"stars"`
}

func Marshall(movie model.Movie) interface{} {
	movieJson, errUn := json.Marshal(movie)
	fmt.Println(errUn)

	var movieM MovieJson
	_ = json.Unmarshal(movieJson, &movieM)
	return movieM
}

func MarshallArray(movies []model.Movie) []interface{} {
	result := make([]interface{}, len(movies))
	for index, movie := range movies {
		result[index] = Marshall(movie)
	}
	return result
}
