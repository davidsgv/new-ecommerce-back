package rest

// import (
// 	"net/http"
// 	"seguridad/core/service"
// 	"strconv"

// 	"github.com/ErnestoGonzalezVargas/rest"
// 	"github.com/ErnestoGonzalezVargas/rest/seguridad"
// 	"github.com/ErnestoGonzalezVargas/rest/seguridad/autenticacion"
// 	"github.com/gin-gonic/gin"
// )

// type AutenticacionHandler struct {
// 	servicioAutenticacion *service.AutenticacionServicio
// }

// func NewAutenticacionHandler(mainRouter *gin.Engine, groupRouter *gin.RouterGroup, autenticacionServicio *service.AutenticacionServicio) {
// 	handler := &AutenticacionHandler{
// 		servicioAutenticacion: autenticacionServicio,
// 	}

// 	mainRouter.POST("/api/login/", handler.Login)

// 	//groupRouter.GET("/roles/", handler.GetRoles)
// 	//groupRouter.GET("/roles/:id/", handler.GetRolById)

// 	groupRouter.POST("/bloqueos/", handler.BloquearUsuario)

// 	groupRouter.GET("/permisos/", handler.GetPermisos)

// 	groupRouter.GET("/conexiones/", handler.GetConexiones)
// 	groupRouter.GET("/conexiones/:id/", handler.GetConexionesByUsuarioId)

// 	//solo para pruebas
// 	mainRouter.POST("/api/sesion/", handler.Sesion)
// }

// func (handler *AutenticacionHandler) Login(ctx *gin.Context) {
// 	request := autenticacion.LoginRequest{}
// 	err := ctx.ShouldBind(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, rest.BadBinding, err.Error()))
// 		return
// 	}

// 	token, err := handler.servicioAutenticacion.IniciarSesion(request.Correo, request.Password, ctx.Request.Host, nil)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	responseData := autenticacion.LoginResponse{
// 		Token: *token,
// 	}

// 	response := rest.NewGenericResponse(responseData, 0, "")
// 	ctx.JSON(http.StatusOK, response)
// }

// func (handler *AutenticacionHandler) Sesion(ctx *gin.Context) {
// 	request := autenticacion.SesionRequest{}
// 	err := ctx.ShouldBind(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, rest.BadBinding, err.Error()))
// 		return
// 	}

// 	_, valido, err := handler.servicioAutenticacion.ValidarSesion(request.Token, request.Modulo, request.Operacion, request.Recurso)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	response := autenticacion.SesionResponse{
// 		Permiso: valido,
// 	}
// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(response, 0, ""))
// }

// // func (handler *AutenticacionHandler) GetRoles(ctx *gin.Context) {
// // 	roles, err := handler.servicioAutenticacion.GetRoles()
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// // 		return
// // 	}

// // 	respuesta := []autenticacion.RolByIdResponse{}
// // 	for _, v := range roles {
// // 		index := slices.IndexFunc(respuesta, func(i autenticacion.RolByIdResponse) bool {
// // 			return i.Id == v.Rol.Id
// // 		})

// // 		if index == -1 {
// // 			rol := autenticacion.RolByIdResponse{
// // 				Id:          v.Rol.Id,
// // 				Nombre:      v.Rol.Nombre,
// // 				Descripcion: v.Rol.Descripcion,
// // 				Permisos:    []autenticacion.Permiso{},
// // 			}

// // 			respuesta = append(respuesta, rol)
// // 			index = len(respuesta) - 1
// // 		}

// // 		permiso := autenticacion.Permiso{
// // 			Id:        v.Permiso.Id,
// // 			Modulo:    v.Permiso.Modulo,
// // 			Recurso:   v.Permiso.Recurso,
// // 			Crear:     v.Crear,
// // 			Eliminar:  v.Eliminar,
// // 			Consultar: v.Consultar,
// // 			Editar:    v.Editar,
// // 		}

// // 		respuesta[index].Permisos = append(respuesta[index].Permisos, permiso)
// // 	}

// // 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// // }

// // func (handler *AutenticacionHandler) GetRolById(ctx *gin.Context) {
// // 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// // 		return
// // 	}

// // 	roles, err := handler.servicioAutenticacion.GetRolById(uint(id))
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// // 		return
// // 	}

// // 	respuesta := autenticacion.RolByIdResponse{}
// // 	for _, v := range roles {
// // 		if respuesta.Id == 0 {
// // 			respuesta.Id = uint(v.Rol.Id)
// // 			respuesta.Nombre = v.Rol.Nombre
// // 			respuesta.Descripcion = v.Rol.Descripcion
// // 			respuesta.Permisos = []autenticacion.Permiso{}
// // 		}

// // 		permiso := autenticacion.Permiso{
// // 			Id:        uint(v.Permiso.Id),
// // 			Modulo:    v.Permiso.Modulo,
// // 			Recurso:   v.Permiso.Recurso,
// // 			Crear:     v.Crear,
// // 			Eliminar:  v.Eliminar,
// // 			Consultar: v.Consultar,
// // 			Editar:    v.Editar,
// // 		}

// // 		respuesta.Permisos = append(respuesta.Permisos, permiso)
// // 	}

// // 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// // }

// func (handler *AutenticacionHandler) BloquearUsuario(ctx *gin.Context) {
// 	correo := ctx.GetString("correo")
// 	err := handler.servicioAutenticacion.BloquearSesion(correo)
// 	respuesta := seguridad.GetBloqueoResponse{
// 		Resultado: true,
// 	}
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// }

// func (handler *AutenticacionHandler) GetPermisos(ctx *gin.Context) {
// 	permisos, err := handler.servicioAutenticacion.GetPermisos()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	respuesta := []autenticacion.GetPermisosResponse{}
// 	for _, v := range permisos {
// 		permiso := autenticacion.GetPermisosResponse{
// 			Id:      uint(v.Id),
// 			Modulo:  v.Modulo,
// 			Recurso: v.Recurso,
// 		}

// 		respuesta = append(respuesta, permiso)
// 	}

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// }

// func (handler *AutenticacionHandler) GetConexiones(ctx *gin.Context) {
// 	conexiones, err := handler.servicioAutenticacion.GetConexiones()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	respuesta := []autenticacion.GetConexionesResponse{}
// 	for _, v := range conexiones {
// 		conexion := autenticacion.GetConexionesResponse{
// 			Id:               v.Id,
// 			Token:            v.Token,
// 			Ingreso:          v.Ingreso,
// 			Equipo:           v.Equipo,
// 			VencimientoToken: v.VencimientoToken,
// 			IdUsuario:        v.IdUsuario,
// 		}

// 		respuesta = append(respuesta, conexion)
// 	}

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// }

// func (handler *AutenticacionHandler) GetConexionesByUsuarioId(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	conexiones, err := handler.servicioAutenticacion.GetConexionesByUsuarioId(uint(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	respuesta := []autenticacion.GetConexionesResponse{}
// 	for _, v := range conexiones {
// 		conexion := autenticacion.GetConexionesResponse{
// 			Id:               v.Id,
// 			Token:            v.Token,
// 			Ingreso:          v.Ingreso,
// 			Equipo:           v.Equipo,
// 			VencimientoToken: v.VencimientoToken,
// 			IdUsuario:        v.IdUsuario,
// 		}

// 		respuesta = append(respuesta, conexion)
// 	}

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// }
