package controller

import (
	"net/http"

	"github.com/Piyawat-T/go-service-server/bootstrap"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	Env *bootstrap.Env
}

func (h *HomeController) Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
