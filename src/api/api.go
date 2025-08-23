package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky6890/golang-clean-web-api/api/routers"
	"github.com/lucky6890/golang-clean-web-api/api/validations"
	"github.com/lucky6890/golang-clean-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.MobileNumberValidator, true)
	}
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
