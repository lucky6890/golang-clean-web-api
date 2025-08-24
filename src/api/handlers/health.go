package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucky6890/golang-clean-web-api/api/helpers"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.GenerateBaseResponse("working", true, 0))
}
