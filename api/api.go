package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		a := v1.Group("trivia")
		{
			fmt.Print(a)
		}
	}
	r.GET("api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
