package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machado-br/k8s-api/api/payloads"
	"github.com/machado-br/k8s-api/models"
)

func (a API) retrieveAll(c *gin.Context) {

	result, err := a.Service.RetrieveAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (a API) create(c *gin.Context) {

	var payload payloads.Namespace
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	namespace := models.Namespace{
		Name: payload.Name,
	}

	err = a.Service.CreateNamespace(c, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (a API) delete(c *gin.Context) {

	name := c.Param("name")

	err := a.Service.DeleteNamespace(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (a API) retrieve(c *gin.Context) {

	name := c.Param("name")

	result, err := a.Service.Retrieve(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
