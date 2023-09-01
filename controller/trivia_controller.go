package controller

import  "github.com/gin-gonic/gin"

type TriviaController interface {
	ObtenerPreguntas(ctx gin.Context)
	AgregarPreguntas(ctx gin.Context, model.Pregunta)
	ModificarPregunta(ctx gin.Context, model.Pregunta)
}

