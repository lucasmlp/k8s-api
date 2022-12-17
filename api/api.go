package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
}

func NewApi(
) (api, error) {
	return api{
	}, nil
}

func (a api) Engine() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	root := router.Group("")
	{
		root.GET("/ping", func(c *gin.Context) {
			log.Printf("ClientIP: %s\n", c.ClientIP())

			c.JSON(http.StatusOK, "pong")
		})
	}

	return router
}

func (a api) Run() {

	router := a.Engine()
	router.Run()
}
