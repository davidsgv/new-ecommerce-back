package repository

import (
	"seguridad/core/domain"
	"time"
)

type iusuario interface {
	BloquearSesion(string) error
	GetUsuarioLogin(correo, dominio string) (*domain.Usuario, error)
	GetRolesByUsuarioId(id uint) ([]domain.Rol, error)
	ValidarUsuario(correo, dominio string, fechaFirma time.Time) (bool, error)
	ValidarPermisos(recurso, modulo, operacion string, roles []string) (bool, error)
}

type iconexion interface {
	SaveConexion(*domain.Conexion) error
	GetConexiones() ([]domain.Conexion, error)
	GetConexionesByUsuarioId(uint) ([]domain.Conexion, error)
}

type ipermisos interface {
	GetRoles() ([]domain.PermisosPorRol, error)
	GetRolById(uint) ([]domain.PermisosPorRol, error)
	GetPermisos() ([]domain.Permiso, error)
}

type IRepositoryAutenticacion interface {
	iusuario
	iconexion
	ipermisos
	//GetEmpresasByUsuarioId(id uint, dominio string) ([]domain.Empresa, error)//se elimina por que GetUsuarioLogin ya valida el dominio
}
