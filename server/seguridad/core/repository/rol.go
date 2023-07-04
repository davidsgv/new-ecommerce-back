package repository

import "seguridad/core/domain"

type iPermiso interface {
	GetPermisos() ([]domain.Permiso, error)
	//GetPermisoById() (*domain.Permiso, error)
}

type iRol interface {
	//operaciones BD
	CreateRol(domain.Rol) (insertedId int64, err error)
	UpdateRol(domain.Rol) error
	GetRoles() ([]domain.Rol, error)
	GetRolByRolId(int64) (*domain.Rol, error)
	DeleteRol(rolId int64) error
	AddPermiso(rolId, permisoId int64) error
	RemovePermiso(rolId, permisoId int64) error

	//para validaciones del servicio
	ExistRol(name string) (bool, error)
}

type IRepositoryRol interface {
	iPermiso
	iRol
}
