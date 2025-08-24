package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky6890/golang-clean-web-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}
