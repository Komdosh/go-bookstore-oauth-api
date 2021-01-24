module github.com/Komdosh/go-bookstore-oauth-api/src

go 1.15

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

require (
	github.com/Komdosh/go-bookstore-utils v0.0.0-20210116090756-a7d3cfa03af1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty v1.12.0
	github.com/gocql/gocql v0.0.0-20201215165327-e49edf966d90
)
