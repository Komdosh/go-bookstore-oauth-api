module github.com/Komdosh/go-bookstore-oauth-api

go 1.15

require (
	github.com/Komdosh/go-bookstore-users-api v0.0.0-20210115060207-219326d2a46f
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty v1.12.0
	github.com/gocql/gocql v0.0.0-20201215165327-e49edf966d90
	github.com/stretchr/testify v1.7.0
	gopkg.in/resty.v1 v1.12.0 // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
