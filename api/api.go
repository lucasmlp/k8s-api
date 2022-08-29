package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/k8s-api/services/listReleases"
)

type api struct {
	ListReleasesService   listReleases.Service
}

func NewApi(
	ListReleasesService   listReleases.Service,
) (api, error) {
	return api{
		ListReleasesService: ListReleasesService,
	}, nil
}

func (a api) Engine() *gin.Engine {
	router := gin.Default()

	root := router.Group("")
	{
		root.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}
	helmRoot := router.Group("/helm")
	{
		helmRoot.GET("", a.allReleases)
	}

	return router
}

func (a api) Run() {

	router := a.Engine()
	router.Run()
}
