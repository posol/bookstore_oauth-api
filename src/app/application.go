package app

import (
	"github.com/gin-gonic/gin"
	"github.com/posol/bookstore_oauth-api/src/http"
	"github.com/posol/bookstore_oauth-api/src/repository/db"
	"github.com/posol/bookstore_oauth-api/src/repository/rest"
	"github.com/posol/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/api/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/api/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
