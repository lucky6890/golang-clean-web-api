package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucky6890/golang-clean-web-api/api/handlers"
	"github.com/lucky6890/golang-clean-web-api/api/middlewares"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", middlewares.Middleware(), handler.Health)
}
