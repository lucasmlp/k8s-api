package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/k8s-api/adapters/k8s"
	"github.com/machado-br/k8s-api/services/namespace"
)

type API struct {
	Deployed         bool
	Adapter          k8s.Adapter
	NamespaceService namespace.Service
}

func NewApi(
	adapter k8s.Adapter,
	namespaceService namespace.Service,
) (API, error) {
	return API{
		Adapter:          adapter,
		NamespaceService: namespaceService,
	}, nil
}

func (a API) Engine() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	root := router.Group("")
	{
		root.GET("/ping", func(c *gin.Context) {
			log.Printf("ClientIP: %s\n", c.ClientIP())

			c.JSON(http.StatusOK, "pong")
		})

		root.GET("/", a.retrieveAll)
		root.POST("/", a.create)
		root.DELETE("/:name", a.delete)
	}

	return router
}

func (a API) Run() {

	router := a.Engine()
	router.Run()
}
