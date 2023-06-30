package rest

// import (
// 	"net/http"
// 	"seguridad/core/domain"
// 	"seguridad/core/service"

// 	"github.com/ErnestoGonzalezVargas/rest"
// 	"github.com/ErnestoGonzalezVargas/rest/seguridad"
// 	"github.com/gin-gonic/gin"
// )

// type EmpresasHandler struct {
// 	servicioEmpresa *service.EmpresaServicio
// }

// func NewEmpresasHandler(groupRouter *gin.RouterGroup, servicioEmpresa *service.EmpresaServicio) {
// 	handler := &EmpresasHandler{
// 		servicioEmpresa: servicioEmpresa,
// 	}

// 	groupRouter.POST("/servidores/", handler.CreateServidor)
// 	groupRouter.PUT("/servidores/", handler.UpdateServidor)
// 	groupRouter.GET("/servidores/", handler.GetServidores)
// }

// func (handler *EmpresasHandler) CreateServidor(ctx *gin.Context) {
// 	request := seguridad.CreateServidorRequest{}
// 	response := seguridad.CreateServidorResponse{}
// 	err := ctx.ShouldBind(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	servidor := domain.Servidor{
// 		Dominio:     request.Dominio,
// 		DireccionIp: request.DireccionIp,
// 		NombreBD:    request.NombreBD,
// 		Usuario:     request.Usuario,
// 		Password:    request.Password,
// 		IdEmpresa:   request.IdEmpresa,
// 	}
// 	servidorCreado, err := handler.servicioEmpresa.CreateServidor(&servidor)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	response.Id = servidorCreado.Id
// 	response.Dominio = servidorCreado.Dominio
// 	response.IdEmpresa = servidorCreado.IdEmpresa

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(response, 0, ""))
// }

// func (handler *EmpresasHandler) UpdateServidor(ctx *gin.Context) {
// 	request := seguridad.UpdateServidorRequest{}
// 	response := seguridad.UpdateServidorResponse{}
// 	err := ctx.ShouldBind(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	servidor := domain.Servidor{
// 		Id:          request.Id,
// 		Dominio:     request.Dominio,
// 		DireccionIp: request.DireccionIp,
// 		NombreBD:    request.NombreBD,
// 		Usuario:     request.Usuario,
// 		Password:    request.Password,
// 		IdEmpresa:   request.IdEmpresa,
// 	}
// 	servidorActualizado, err := handler.servicioEmpresa.UpdateServidor(&servidor)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	response.Id = servidorActualizado.Id
// 	response.Dominio = servidorActualizado.Dominio
// 	response.IdEmpresa = servidorActualizado.IdEmpresa

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(response, 0, ""))
// }

// func (handler *EmpresasHandler) GetServidores(ctx *gin.Context) {
// 	servidores, err := handler.servicioEmpresa.GetServidores()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	response := []seguridad.GetServidorResponse{}
// 	for _, v := range servidores {
// 		servidor := seguridad.GetServidorResponse{
// 			Id:        v.Id,
// 			Dominio:   v.Dominio,
// 			EmpresaId: v.IdEmpresa,
// 		}
// 		response = append(response, servidor)
// 	}

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(response, 0, ""))
// }
