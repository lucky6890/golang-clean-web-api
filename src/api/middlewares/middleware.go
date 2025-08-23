package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//do something
		fmt.Print("middleware called")
		ctx.Next()
	}
}
