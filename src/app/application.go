package app

import (
	"github.com/gin-gonic/gin"
	"github.com/posol/bookstore_oauth-api/clients/cassandra"
	"github.com/posol/bookstore_oauth-api/src/domain/access_token"
	"github.com/posol/bookstore_oauth-api/src/http"
	"github.com/posol/bookstore_oauth-api/src/repository/db"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/api/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/api/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
