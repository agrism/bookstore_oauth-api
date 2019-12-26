package app

import (
	"github.com/agrism/bookstore_oauth-api/src/domain/access_token"
	"github.com/agrism/bookstore_oauth-api/src/domain/repository/db"
	"github.com/agrism/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run("127.0.0.1:8081")
}
