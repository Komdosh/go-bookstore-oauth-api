package app

import (
	"github.com/Komdosh/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/Komdosh/go-bookstore-oauth-api/src/http"
	"github.com/Komdosh/go-bookstore-oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
