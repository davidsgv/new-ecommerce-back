package service

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"seguridad/core/domain"
// 	"seguridad/core/repository"
// 	"time"
// 	"validacion"

// 	"github.com/dgrijalva/jwt-go"
// )

// type jwtCustomClaims struct {
// 	Correo string
// 	Roles  []string
// 	jwt.StandardClaims
// }

// type AutenticacionServicio struct {
// 	repo            repository.IRepositoryAutenticacion
// 	tokenExpireTime time.Duration //horas
// 	key             string
// 	validador       *validacion.Validador
// }

// func NewAutenticacionServicio(repo repository.IRepositoryAutenticacion, tokenExpire time.Duration, key string, validador *validacion.Validador) *AutenticacionServicio {
// 	return &AutenticacionServicio{
// 		repo:            repo,
// 		tokenExpireTime: tokenExpire,
// 		key:             key,
// 		validador:       validador,
// 	}
// }

// // region sesion
// func (servicio *AutenticacionServicio) IniciarSesion(correo, password, dominio string, equipo *string) (*string, error) {
// 	err := servicio.validador.Verificar("correo", correo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = servicio.validador.Verificar("dominio", dominio)
// 	if err != nil {
// 		return nil, err
// 	}

// 	now := time.Now()
// 	fechaVencimientoToken := now.Add(servicio.tokenExpireTime * time.Hour)

// 	//busca el usuario en BD
// 	usuario, err := servicio.repo.GetUsuarioLogin(correo, dominio)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//si no trae datos no existe
// 	if usuario == nil {
// 		return nil, errors.New("credenciales invalidas")
// 	}

// 	//cifra la contraseña del usuario
// 	hash, err := cifrarContraseña(password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//compara la contraseña con la guardada en BD
// 	if hash != usuario.Password {
// 		return nil, errors.New("credenciales invalidas")
// 	}

// 	//se firma el token
// 	roles, err := servicio.repo.GetRolesByUsuarioId(usuario.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var rolesClaim = make([]string, 0)
// 	for _, v := range roles {
// 		rolesClaim = append(rolesClaim, v.Nombre)
// 	}

// 	claim := jwtCustomClaims{
// 		Correo: usuario.Correo,
// 		Roles:  rolesClaim,
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer:    dominio,
// 			IssuedAt:  now.Unix(),
// 			ExpiresAt: fechaVencimientoToken.Unix(),
// 		},
// 	}

// 	signer := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)
// 	token, err := signer.SignedString([]byte(servicio.key))
// 	if err != nil {
// 		return nil, err
// 	}

// 	//se guarda la nueva conexion
// 	conexion := &domain.Conexion{
// 		Token:            token,
// 		Ingreso:          now,
// 		Equipo:           equipo,
// 		VencimientoToken: fechaVencimientoToken,
// 		IdUsuario:        usuario.Id,
// 	}
// 	err = servicio.repo.SaveConexion(conexion)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &token, nil
// }

// func (servicio *AutenticacionServicio) ValidarSesion(token, modulo, operacion, recurso string) (*jwtCustomClaims, bool, error) {
// 	tokenValidation, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		// Signing method validation
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			errorMensaje := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
// 			log.Println(errorMensaje)
// 			return nil, errors.New(errorMensaje)
// 		}
// 		// Return the secret signing key
// 		return []byte(servicio.key), nil
// 	})
// 	if err != nil {
// 		return nil, false, err
// 	}

// 	if !tokenValidation.Valid {
// 		return nil, false, errors.New("token invalido")
// 	}

// 	mapClaims := tokenValidation.Claims.(jwt.MapClaims)

// 	roles := []string{}
// 	for _, v := range mapClaims["Roles"].([]interface{}) {
// 		roles = append(roles, v.(string))
// 	}

// 	claims := jwtCustomClaims{
// 		Correo: mapClaims["Correo"].(string),
// 		Roles:  roles,
// 		StandardClaims: jwt.StandardClaims{
// 			Issuer:   mapClaims["iss"].(string),
// 			IssuedAt: int64(mapClaims["iat"].(float64)),
// 		},
// 	}
// 	fechaFirma := time.Unix(claims.IssuedAt, 0)

// 	//valida si esta activo, bloqueos y dominio
// 	valido, err := servicio.repo.ValidarUsuario(claims.Correo, claims.Issuer, fechaFirma)
// 	if err != nil {
// 		return nil, false, err
// 	}

// 	if !valido {
// 		return nil, false, errors.New("usuario restringido")
// 	}

// 	//valida los roles y permisos
// 	tieneAcceso, err := servicio.repo.ValidarPermisos(recurso, modulo, operacion, claims.Roles)
// 	if err != nil {
// 		return nil, false, err
// 	}

// 	if !tieneAcceso {
// 		return nil, false, errors.New("el usuario no cuenta con los permisos necesarios")
// 	}

// 	return &claims, tieneAcceso, nil
// }

// func (servicio *AutenticacionServicio) BloquearSesion(correo string) error {
// 	servicio.validador.Verificar("email", correo)
// 	err := servicio.repo.BloquearSesion(correo)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //endregion

// // region roles
// func (servicio *AutenticacionServicio) GetRoles() ([]domain.PermisosPorRol, error) {
// 	permisosRol, err := servicio.repo.GetRoles()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if permisosRol == nil {
// 		return nil, errors.New("no se encontraron roles")
// 	}

// 	return permisosRol, nil
// }

// func (servicio *AutenticacionServicio) GetRolById(id uint) ([]domain.PermisosPorRol, error) {
// 	err := servicio.validador.Verificar("id", id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	permisosRol, err := servicio.repo.GetRolById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if permisosRol == nil {
// 		return nil, errors.New("no se encontraron roles")
// 	}

// 	return permisosRol, nil
// }

// //endregion

// func (servicio *AutenticacionServicio) GetPermisos() ([]domain.Permiso, error) {
// 	permisos, err := servicio.repo.GetPermisos()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return permisos, nil
// }

// func (servicio *AutenticacionServicio) GetConexiones() ([]domain.Conexion, error) {
// 	conexiones, err := servicio.repo.GetConexiones()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return conexiones, nil
// }

// func (servicio *AutenticacionServicio) GetConexionesByUsuarioId(usuarioId uint) ([]domain.Conexion, error) {
// 	err := servicio.validador.Verificar("id", usuarioId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	conexiones, err := servicio.repo.GetConexionesByUsuarioId(usuarioId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return conexiones, nil
// }

// //pendiente
// /*
// //obtener conexiones
// 	//devuelve lista de conexiones

// //cerrar sesion todos dispositivos

// //desactivar usuario
// 	//validar dominio y empresa
// 	//bloquear usuario

// //crear usuario
// */
