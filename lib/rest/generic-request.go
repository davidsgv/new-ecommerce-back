package rest

import (
	"net/http"
	"respuesta"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Codigo  uint    `json:"codigo"`
	Mensaje *string `json:"mensaje"`
}

type GenericResponse struct {
	Error ErrorResponse `json:"error"`
	Data  interface{}   `json:"data"`
}

func ResponseBadRequest(ctx *gin.Context, err error) {
	var mensajeError *string
	if err != nil {
		mes := err.Error()
		mensajeError = &mes
	}

	response := &GenericResponse{
		Error: ErrorResponse{
			Codigo:  respuesta.BadBinding,
			Mensaje: mensajeError,
		},
		Data: nil,
	}

	ctx.JSON(http.StatusBadRequest, response)
}

func NewGenericResponse(data interface{}, errorCode uint, mensaje *string) *GenericResponse {
	return &GenericResponse{
		Error: ErrorResponse{
			Codigo:  errorCode,
			Mensaje: mensaje,
		},
		Data: data,
	}
}

func Response(ctx *gin.Context, resSer *respuesta.Respuesta, restData interface{}) {
	var httpCode int
	var mensaje string

	switch {
	case resSer.Codigo > 0 && resSer.Codigo < 10 || resSer.Codigo == 201:
		httpCode = http.StatusOK
		mensaje = resSer.Error.Error()
	case resSer.Codigo == 10:
		httpCode = http.StatusInternalServerError
		mensaje = "Error interno"
	case resSer.Codigo > 200 && resSer.Codigo < 299:
		httpCode = http.StatusBadRequest
		mensaje = resSer.Error.Error()
	}

	info := &mensaje
	if mensaje == "" {
		info = nil
	}

	response := GenericResponse{
		Error: ErrorResponse{
			Codigo:  uint(resSer.Codigo),
			Mensaje: info,
		},
		Data: restData,
	}
	ctx.JSON(httpCode, response)
}
