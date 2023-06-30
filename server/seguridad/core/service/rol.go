package service

import (
	"fmt"
	"respuesta"
	"seguridad/core/domain"
	"seguridad/core/repository"
	"validacion"
)

type RolService struct {
	repo      repository.IRepositoryRol
	validator *validacion.Validador
}

func NewRolService(repo repository.IRepositoryRol, val *validacion.Validador) *RolService {
	return &RolService{
		repo:      repo,
		validator: val,
	}
}

func (servicio *RolService) CreateRol(rol domain.Rol) *respuesta.Respuesta {
	//validar datos
	err := servicio.validator.Verificar("min=6,max=200", rol.Nombre)
	if err != nil {
		err = fmt.Errorf("Nombre: %w", err)
		return respuesta.NewErrValidation(err)
	}

	err = servicio.validator.Verificar("min=50,max=2000", rol.Descripcion)
	if err != nil {
		err = fmt.Errorf("Descripción: %w", err)
		return respuesta.NewErrValidation(err)
	}

	//revisar si el registro no existe
	exists, err := servicio.repo.ExistRol(rol.Nombre)
	if err != nil {
		return respuesta.NewErrRepository(err)
	}

	if exists {
		err = fmt.Errorf("El rol %s ya existe", rol.Nombre)
		return respuesta.NewErrDuplicateRow(err)
	}

	//crear el registro
	id, err := servicio.repo.CreateRol(rol)
	if err != nil {
		return respuesta.NewErrRepository(err)
	}

	rol.Id = id
	return respuesta.NewRespuesta(respuesta.NoError, nil, &rol)
}

func (servicio *RolService) UpdateRol(rol domain.Rol) *respuesta.Respuesta {
	//validar datos
	err := servicio.validator.Verificar("id", rol.Id)
	if err != nil {
		return respuesta.NewErrValidation(err)
	}

	err = servicio.validator.Verificar("min=6,max=200", rol.Nombre)
	if err != nil {
		return respuesta.NewErrValidation(err)
	}

	err = servicio.validator.Verificar("min=50,max=2000", rol.Descripcion)
	if err != nil {
		return respuesta.NewErrValidation(err)
	}

	//revisar si el rol existe
	exists, err := servicio.repo.ExistRol(rol.Nombre)
	if err != nil {
		return respuesta.NewErrRepository(err)
	}

	if exists {
		err = fmt.Errorf("El rol %s ya existe", rol.Nombre)
		return respuesta.NewErrDuplicateRow(err)
	}

	//actualizar el registro
	err = servicio.repo.UpdateRol(rol)
	if err != nil {
		return respuesta.NewErrRepository(err)
	}

	return respuesta.NewNoErr(&rol)
}

func (servicio *RolService) GetRoles() *respuesta.Respuesta {
	roles, err := servicio.repo.GetRoles()
	if err != nil {
		return respuesta.NewErrRepository(err)
	}
	if roles == nil {
		return respuesta.NewNoData()
	}
	return respuesta.NewNoErr(roles)
}

func (servicio *RolService) GetRolByRolId(rolId int64) *respuesta.Respuesta {
	//validar datos
	err := servicio.validator.Verificar("id", rolId)
	if err != nil {
		return respuesta.NewErrValidation(err)
	}

	//buscar en el repositorio
	permisosRol, err := servicio.repo.GetRolByRolId(rolId)
	if err != nil {
		return respuesta.NewErrRepository(err)
	}

	if permisosRol == nil {
		return respuesta.NewNoData()
	}

	return respuesta.NewNoErr(permisosRol)

}

// func (servicio *RolService) AddPermiso(rolId, permisoId int64) *respuesta.Respuesta {
// 	//validar datos
// 	err := servicio.validator.Verificar("id", rolId)
// 	if err != nil {
// 		return respuesta.NewErrValidation(err)
// 	}
// 	err = servicio.validator.Verificar("id", permisoId)
// 	if err != nil {
// 		return respuesta.NewErrValidation(err)
// 	}

// 	//insertar permiso
// 	err = servicio.repo.AddPermiso(rolId, permisoId)
// 	if err != nil {
// 		return respuesta.NewErrRepository(err)
// 	}

// 	resRoles := servicio.GetRolByRolId(rolId)
// 	if resRoles.Codigo != respuesta.NoError {
// 		return respuesta.NewRespuesta(respuesta.RegistroCreadoErrorObteniendoDatos, resRoles.Error, nil)
// 	}
// 	return respuesta.NewNoErr(&resRoles.Datos)
// }

func (servicio *RolService) RemovePermiso(rolId, permisoId int64) error {
	return fmt.Errorf("Not Implemented Yet")
}

func (servicio *RolService) DeleteRol(rolId int64) error {
	return fmt.Errorf("Not Implemented Yet")
}
