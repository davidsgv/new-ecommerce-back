package service_test

import (
	"errors"
	"seguridad/core/domain"
)

var datosRoles = []domain.Rol{
	{
		Id:          1,
		Nombre:      "administrador",
		Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
		Permisos: []domain.PermisosPorRol{
			{
				Crear:     true,
				Editar:    true,
				Consultar: true,
				Eliminar:  true,
				Permiso: domain.Permiso{
					Id:      1,
					Modulo:  "seguridad",
					Recurso: "usuarios",
				},
			},
			{
				Crear:     true,
				Editar:    true,
				Consultar: true,
				Eliminar:  true,
				Permiso: domain.Permiso{
					Id:      1,
					Modulo:  "seguridad",
					Recurso: "roles",
				},
			},
		},
	},
	{
		Id:          2,
		Nombre:      "Asesor",
		Descripcion: "Este rol tiene privilegios reducidos, puede hacer operaciones limitadas en cuanto a escribir datos",
		Permisos: []domain.PermisosPorRol{
			{
				Crear:     false,
				Editar:    false,
				Consultar: false,
				Eliminar:  false,
				Permiso: domain.Permiso{
					Id:      1,
					Modulo:  "seguridad",
					Recurso: "usuarios",
				},
			},
			{
				Crear:     false,
				Editar:    false,
				Consultar: true,
				Eliminar:  false,
				Permiso: domain.Permiso{
					Id:      1,
					Modulo:  "seguridad",
					Recurso: "roles",
				},
			},
		},
	},
}

// region succes mock
type mockedCreateRolSucces struct {
}

func (mock *mockedCreateRolSucces) CreateRol(domain.Rol) (insertedId int64, err error) {
	return 1, nil
}

type mockedUpdateRolSucces struct {
}

func (mock *mockedUpdateRolSucces) UpdateRol(domain.Rol) error {
	return nil
}

type mockedGetRolesSucces struct {
}

func (mock *mockedGetRolesSucces) GetRoles() ([]domain.Rol, error) {
	return datosRoles, nil
}

type mockedGetRolByRolIdSucces struct {
}

func (mock *mockedGetRolByRolIdSucces) GetRolByRolId(id int64) (*domain.Rol, error) {
	for _, v := range datosRoles {
		if v.Id == id {
			return &v, nil
		}
	}
	return &datosRoles[0], nil
}

type mockedDeleteRolSucces struct {
}

func (mock *mockedDeleteRolSucces) DeleteRol(rolId int64) error {
	return nil
}

type mockedExistRolSucces struct {
}

func (mock *mockedExistRolSucces) ExistRol(name string) (bool, error) {
	return false, nil
}

type mockedAddPermisoSucces struct {
}

func (mock *mockedAddPermisoSucces) AddPermiso(rolId, permisoId int64) error {
	return nil
}

type mockedRemovePermisoSucces struct {
}

func (mock *mockedRemovePermisoSucces) RemovePermiso(rolId, permisoId int64) error {
	return nil
}

type mockedGetPermisosSucces struct {
}

func (mock *mockedGetPermisosSucces) GetPermisos() ([]domain.Permiso, error) {
	return nil, nil
}

//endregion

// region fail mock
type mockedCreateRolFail struct {
}

func (mock *mockedCreateRolFail) CreateRol(domain.Rol) (insertedId int64, err error) {
	return 0, errors.New("Falla al crear el rol")
}

type mockedUpdateRolFail struct {
}

func (mock *mockedUpdateRolFail) UpdateRol(domain.Rol) error {
	return errors.New("Falla al actualizar el rol")
}

type mockedGetRolesFail struct {
}

func (mock *mockedGetRolesFail) GetRoles() ([]domain.Rol, error) {
	return nil, errors.New("Falla al obtener los roles")
}

type mockedGetRolByRolIdFail struct {
}

func (mock *mockedGetRolByRolIdFail) GetRolByRolId(int64) (*domain.Rol, error) {
	return nil, errors.New("Falla al obtener el rol")
}

type mockedDeleteRolFail struct {
}

func (mock *mockedDeleteRolFail) DeleteRol(rolId int64) error {
	return errors.New("Falla al eliminar el rol")
}

type mockedExistRolFail struct {
}

func (mock *mockedExistRolFail) ExistRol(name string) (bool, error) {
	return false, errors.New("Falla al verificar el rol")
}

type mockedAddPermisoFail struct {
}

func (mock *mockedAddPermisoFail) AddPermiso(rolId, permisoId int64) error {
	return errors.New("Falla al crear el permiso")
}

type mockedRemovePermisoFail struct {
}

func (mock *mockedRemovePermisoFail) RemovePermiso(rolId, permisoId int64) error {
	return errors.New("Falla al eliminar el permiso")
}

type mockedGetPermisosFail struct {
}

func (mock *mockedGetPermisosFail) GetPermisos() ([]domain.Permiso, error) {
	return nil, errors.New("Falla al obtener el permiso")
}

//endregion

// region empty mock
type mockedCreateRolEmpty struct {
}

func (mock *mockedCreateRolEmpty) CreateRol(domain.Rol) (insertedId int64, err error) {
	return 1, nil
}

type mockedUpdateRolEmpty struct {
}

func (mock *mockedUpdateRolEmpty) UpdateRol(domain.Rol) error {
	return nil
}

type mockedGetRolesEmpty struct {
}

func (mock *mockedGetRolesEmpty) GetRoles() ([]domain.Rol, error) {
	return nil, nil
}

type mockedGetRolByRolIdEmpty struct {
}

func (mock *mockedGetRolByRolIdEmpty) GetRolByRolId(int64) (*domain.Rol, error) {
	return nil, nil
}

type mockedDeleteRolEmpty struct {
}

func (mock *mockedDeleteRolEmpty) DeleteRol(rolId int64) error {
	return nil
}

type mockedExistRolEmpty struct {
}

func (mock *mockedExistRolEmpty) ExistRol(name string) (bool, error) {
	return false, nil
}

type mockedAddPermisoEmpty struct {
}

func (mock *mockedAddPermisoEmpty) AddPermiso(rolId, permisoId int64) error {
	return nil
}

type mockedRemovePermisoEmpty struct {
}

func (mock *mockedRemovePermisoEmpty) RemovePermiso(rolId, permisoId int64) error {
	return nil
}

type mockedGetPermisosEmpty struct {
}

func (mock *mockedGetPermisosEmpty) GetPermisos() ([]domain.Permiso, error) {
	return nil, nil
}

//endregion

type MockedRepositoryRolSucces struct {
	mockedCreateRolSucces
	mockedUpdateRolSucces
	mockedGetRolesSucces
	mockedGetRolByRolIdSucces
	mockedDeleteRolSucces
	mockedExistRolSucces
	mockedAddPermisoSucces
	mockedRemovePermisoSucces
	mockedGetPermisosSucces
}

type MockedRepositoryRolFail struct {
	mockedCreateRolFail
	mockedUpdateRolFail
	mockedGetRolesFail
	mockedGetRolByRolIdFail
	mockedDeleteRolFail
	mockedExistRolFail
	mockedAddPermisoFail
	mockedRemovePermisoFail
	mockedGetPermisosFail
}

type MockedRepositoryRolEmpty struct {
	mockedCreateRolEmpty
	mockedUpdateRolEmpty
	mockedGetRolesEmpty
	mockedGetRolByRolIdEmpty
	mockedDeleteRolEmpty
	mockedExistRolEmpty
	mockedAddPermisoEmpty
	mockedRemovePermisoEmpty
	mockedGetPermisosEmpty
}

type MockedRepositoryCreateRolError struct {
	mockedCreateRolFail
	mockedUpdateRolSucces
	mockedGetRolesSucces
	mockedGetRolByRolIdSucces
	mockedDeleteRolSucces
	mockedExistRolEmpty
	mockedAddPermisoSucces
	mockedRemovePermisoSucces
	mockedGetPermisosSucces
}
type MockedRepositoryRolDuplicateError struct {
	mockedCreateRolSucces
	mockedUpdateRolSucces
	mockedGetRolesSucces
	mockedGetRolByRolIdSucces
	mockedDeleteRolSucces
	mockedAddPermisoSucces
	mockedRemovePermisoSucces
	mockedGetPermisosSucces
}

func (mocked *MockedRepositoryRolDuplicateError) ExistRol(name string) (bool, error) {
	return true, nil
}
