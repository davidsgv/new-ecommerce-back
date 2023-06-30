package rest

// import (
// 	"log"
// 	"net/http"
// 	"seguridad/core/service"
// 	"strings"

// 	"github.com/ErnestoGonzalezVargas/rest"
// 	"github.com/gin-gonic/gin"
// )

// func AuthorizeJWT(servicio *service.AutenticacionServicio) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		token := ctx.GetHeader("Authorization")
// 		operacion := ctx.Request.Method

// 		modulo := "Seguridad"         //ctx.Request.URL.RawPath //Seguridad
// 		uri := ctx.Request.RequestURI //segundo slash
// 		paths := strings.Split(uri, "/")
// 		if len(paths) < 3 {
// 			log.Println("La ruta de la url contiene un path menor a dos separadores (/)")
// 			ctx.AbortWithStatus(http.StatusInternalServerError)
// 			return
// 		}
// 		recurso := paths[2]

// 		claims, tienePermiso, err := servicio.ValidarSesion(token, modulo, operacion, recurso)
// 		if tienePermiso {
// 			ctx.Set("correo", claims.Correo)
// 			var roles string
// 			for i := 0; i < len(claims.Roles); i++ {
// 				roles += claims.Roles[i]
// 				if i != len(claims.Roles) {
// 					roles += ","
// 				}
// 			}
// 			ctx.Set("roles", roles)
// 			return
// 		}

// 		var status int = http.StatusUnauthorized
// 		if err != nil {
// 			switch err.Error() {
// 			// case "el usuario no cuenta con los permisos necesarios":
// 			// case "Token is expired":
// 			// 	status = http.StatusInternalServerError
// 			case "token contains an invalid number of segments":
// 				status = http.StatusBadRequest
// 			default:
// 				status = http.StatusUnauthorized
// 			}
// 		}

// 		response := rest.NewGenericResponse(nil, 0, err.Error())
// 		ctx.AbortWithStatusJSON(status, response)
// 	}
// }
