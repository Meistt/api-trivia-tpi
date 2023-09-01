package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine, controller controller.TriviaController) {
	v1 := r.Group("/api/v1")
	{
		trivia := v1.Group("trivia")
		{
			trivia.GET("", controller.ObtenerPreguntas)
		}
	}
	r.GET("api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
