module github.com/ederj98/movies-microservice

go 1.15

require (
	github.com/fmcarrero/bookstore_utils-go v0.0.0-20200312163057-dc81d1caa362
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-resty/resty/v2 v2.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1
	github.com/testcontainers/testcontainers-go v0.9.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)
replace golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a