package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lucky6890/golang-clean-web-api/api/middlewares"
	"github.com/lucky6890/golang-clean-web-api/api/routers"
	"github.com/lucky6890/golang-clean-web-api/api/validations"
	"github.com/lucky6890/golang-clean-web-api/config"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery())

	registerValidators()

	registerRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}

func registerValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validations.PasswordValidator, true)
	}
}

func registerRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
}
