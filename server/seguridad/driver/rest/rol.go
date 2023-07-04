package rest

import (
	"seguridad/core/domain"
	"seguridad/core/service"

	"github.com/ErnestoGonzalezVargas/rest"
	"github.com/ErnestoGonzalezVargas/rest/seguridad"
	"github.com/gin-gonic/gin"
)

type RolHandler struct {
	service *service.RolService
}

func NewRolHandler(router *gin.RouterGroup, rolServicio *service.RolService) {
	handler := RolHandler{
		service: rolServicio,
	}

	//permisos
	//router.GET("/permisos/", handler.GetPermisos)

	//roles
	router.POST("/roles", handler.CreateRol)
	//router.POST("/roles/permiso/", handler.AddPermiso)

	//router.PUT("/roles/:id", handler.UpdateRol)
	// router.PUT("/roles/permiso", handler.RemovePermiso)

	router.GET("/roles/", handler.GetRoles)
	// router.GET("/roles/:id", handler.GetRolById)

	// router.DELETE("/roles/:id", handler.DeleteRol)

}

func (handler *RolHandler) CreateRol(ctx *gin.Context) {
	//inicializo request
	request := seguridad.CreateRolRequest{}

	//reviso el binding
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		rest.ResponseBadRequest(ctx, err)
		return
	}

	//convierto a entity los datos de entrar
	rol := domain.Rol{
		Nombre:      request.Nombre,
		Descripcion: request.Descripcion,
	}

	//llamo al servicio
	respuesta := handler.service.CreateRol(rol)

	//transformo la respuesta del servicio en response
	var data *domain.Rol
	var response seguridad.CreateRolResponse
	if respuesta.Datos != nil {
		data = respuesta.Datos.(*domain.Rol)
		response = seguridad.CreateRolResponse{
			Id:          data.Id,
			Nombre:      data.Nombre,
			Descripcion: data.Descripcion,
		}
	}

	rest.Response(ctx, respuesta, response)
}

// func (handler *RolHandler) UpdateRol(ctx *gin.Context) {
// 	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
// 	if err != nil {
// 		rest.ResponseBadRequest(ctx, err)
// 		return
// 	}

// 	request := seguridad.CreateRolRequest{}
// 	err = ctx.ShouldBindJSON(&request)
// 	if err != nil {
// 		rest.ResponseBadRequest(ctx, err)
// 		return
// 	}

// 	rol := domain.Rol{
// 		Id:          id,
// 		Nombre:      request.Nombre,
// 		Descripcion: request.Descripcion,
// 	}

// 	//llamo al servicio
// 	respuesta := handler.service.UpdateRol(rol)

// 	//transformo la respuesta del servicio en response
// 	var data domain.Rol
// 	var response seguridad.CreateRolResponse
// 	if respuesta.Datos != nil {
// 		data = respuesta.Datos.(domain.Rol)
// 		response = seguridad.CreateRolResponse{
// 			Id:          data.Id,
// 			Nombre:      data.Nombre,
// 			Descripcion: data.Descripcion,
// 		}
// 	}

// 	rest.Response(ctx, respuesta, response)
// }

func (handler *RolHandler) GetRoles(ctx *gin.Context) {
	respuesta := handler.service.GetRoles()

	//transformo la respuesta del servicio en response
	var data []domain.Rol
	var response []seguridad.GetRolesResponse
	if respuesta.Datos != nil {
		data = respuesta.Datos.([]domain.Rol)

		for _, v := range data {
			res := seguridad.GetRolesResponse{
				Id:          v.Id,
				Nombre:      v.Nombre,
				Descripcion: v.Descripcion,
				Permisos:    []seguridad.GetRolesResponsePermisos{},
			}

			for _, h := range v.Permisos {
				permiso := seguridad.GetRolesResponsePermisos{
					Crear:     h.Crear,
					Editar:    h.Editar,
					Consultar: h.Consultar,
					Eliminar:  h.Eliminar,
					Id:        h.Permiso.Id,
					Modulo:    h.Permiso.Modulo,
					Recurso:   h.Permiso.Recurso,
				}
				res.Permisos = append(res.Permisos, permiso)
			}

			response = append(response, res)
		}
	}

	rest.Response(ctx, respuesta, response)
}

// func arrayContains[T comparable](arr []T, element T) bool {
// 	for _, x := range arr {
// 		if x == element {
// 			return true
// 		}
// 	}
// }
