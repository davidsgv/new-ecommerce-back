package rest

// import (
// 	"net/http"
// 	//serRes "respuestaServicio"
// 	"seguridad/core/domain"
// 	"seguridad/core/service"
// 	"strconv"

// 	"github.com/ErnestoGonzalezVargas/rest"
// 	"github.com/ErnestoGonzalezVargas/rest/seguridad"
// 	"github.com/gin-gonic/gin"
// )

// type UsuarioHandler struct {
// 	servicio *service.UsuarioServicio
// }

// func NewUsuarioHandler(router *gin.RouterGroup, usuarioServicio *service.UsuarioServicio) {
// 	handler := UsuarioHandler{
// 		servicio: usuarioServicio,
// 	}

// 	router.GET("/usuarios/", handler.getUsuario)
// 	router.GET("/usuarios/:id", handler.getUsuarioByID)
// 	//router.POST("/usuarios/", handler.CreateUsuario)
// 	router.PUT("/usuarios/", handler.UpdateUsuario)
// }

// func (handler *UsuarioHandler) getUsuario(ctx *gin.Context) {
// 	usuarios, err := handler.servicio.GetUsuarios()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	respuesta := []seguridad.GetUsuariosResponse{}
// 	for _, v := range usuarios {
// 		usuario := seguridad.GetUsuariosResponse{
// 			Id:        v.Id,
// 			Correo:    v.Correo,
// 			Telefono:  v.Telefono,
// 			Celular:   v.Celular,
// 			Direccion: v.Direccion,
// 		}

// 		respuesta = append(respuesta, usuario)
// 	}

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// }

// func (handler *UsuarioHandler) getUsuarioByID(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	usuarios, err := handler.servicio.GetUsuariosById(uint(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	respuesta := []seguridad.GetUsuariosResponse{}
// 	for _, v := range usuarios {
// 		usuario := seguridad.GetUsuariosResponse{
// 			Id:        v.Id,
// 			Correo:    v.Correo,
// 			Telefono:  v.Telefono,
// 			Celular:   v.Celular,
// 			Direccion: v.Direccion,
// 		}

// 		respuesta = append(respuesta, usuario)
// 	}

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(respuesta, 0, ""))
// }

// // func (handler *UsuarioHandler) CreateUsuario(ctx *gin.Context) {
// // 	request := seguridad.CreateUsuariosRequest{}
// // 	//response := seguridad.GetUsuariosResponse{}

// // 	//err := ctx.ShouldBind(&request)
// // 	// if err != nil {
// // 	// 	res := serRes.NewRespuesta(serRes.ValidacionDatos, err, nil)
// // 	// 	rest.Response(ctx, res, nil)
// // 	// 	return
// // 	// }

// // 	// usuario := domain.Usuario{
// // 	// 	Password:  request.Password,
// // 	// 	Correo:    request.Correo,
// // 	// 	Telefono:  request.Telefono,
// // 	// 	Celular:   request.Celular,
// // 	// 	Direccion: request.Direccion,
// // 	// }
// // 	// servicioRespuesta := handler.servicio.CreateUsuario(usuario)
// // 	// if servicioRespuesta.Codigo != serRes.NoError {
// // 	// 	rest.Response(ctx, servicioRespuesta, nil)
// // 	// 	return
// // 	// }

// // 	// usuarioCreado := servicioRespuesta.Datos.(domain.Usuario)
// // 	// response.Id = usuarioCreado.Id
// // 	// response.Celular = usuarioCreado.Celular
// // 	// response.Correo = usuarioCreado.Correo
// // 	// response.Direccion = usuarioCreado.Direccion
// // 	// response.Telefono = usuarioCreado.Telefono

// // 	// rest.Response(ctx, servicioRespuesta, response)
// // }

// func (handler *UsuarioHandler) UpdateUsuario(ctx *gin.Context) {
// 	request := seguridad.UpdateUsuariosRequest{}
// 	response := seguridad.GetUsuariosResponse{}

// 	err := ctx.ShouldBind(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	usuario := domain.Usuario{
// 		Id:        request.Id,
// 		Correo:    request.Correo,
// 		Telefono:  request.Telefono,
// 		Celular:   request.Celular,
// 		Direccion: request.Direccion,
// 	}
// 	usuarioCreado, err := handler.servicio.UpdateUsuario(usuario)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, rest.NewGenericResponse(nil, 0, err.Error()))
// 		return
// 	}

// 	response.Id = usuarioCreado.Id
// 	response.Celular = usuarioCreado.Celular
// 	response.Correo = usuarioCreado.Correo
// 	response.Direccion = usuarioCreado.Direccion
// 	response.Telefono = usuarioCreado.Telefono

// 	ctx.JSON(http.StatusOK, rest.NewGenericResponse(response, 0, ""))
// }
