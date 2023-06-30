package service_test

import (
	"errors"
	"respuesta"
	"seguridad/core/domain"
	"seguridad/core/repository"
	"seguridad/core/service"
	"testing"
	"validacion"
)

// region test structs
type testRol struct {
	name            string
	rol             domain.Rol
	expectedResult  *respuesta.Respuesta
	mockedInterface repository.IRepositoryRol
}

var datosRoles = []domain.PermisosPorRol{
	{
		Rol: domain.Rol{
			Id:          1,
			Nombre:      "administrador",
			Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
		},
		Crear:     true,
		Editar:    true,
		Consultar: true,
		Eliminar:  true,
		Permiso: domain.Permiso{
			Id:      1,
			Modulo:  "administracion",
			Recurso: "roles",
		},
	},
	{
		Rol: domain.Rol{
			Id:          2,
			Nombre:      "Asesor",
			Descripcion: "Este rol tiene privilegios reducidos, puede hacer operaciones limitadas en cuanto a escribir datos",
		},
		Crear:     false,
		Editar:    false,
		Consultar: true,
		Eliminar:  false,
		Permiso: domain.Permiso{
			Id:      1,
			Modulo:  "administracion",
			Recurso: "roles",
		},
	},
}

//endregion

// mocks para el procedo de testing
// region succes mock
type mockedRepositoryRolSucces struct{}

func (mock *mockedRepositoryRolSucces) CreateRol(domain.Rol) (insertedId int64, err error) {
	return 1, nil
}
func (mock *mockedRepositoryRolSucces) UpdateRol(domain.Rol) error {
	return nil
}
func (mock *mockedRepositoryRolSucces) GetRoles() (*[]domain.PermisosPorRol, error) {
	return &datosRoles, nil
}
func (mock *mockedRepositoryRolSucces) GetRolByRolId(int64) (*[]domain.PermisosPorRol, error) {
	return &datosRoles[0], nil
}
func (mock *mockedRepositoryRolSucces) DeleteRol(rolId int64) error {
	return nil
}
func (mock *mockedRepositoryRolSucces) ExistRol(name string) (bool, error) {
	return false, nil
}
func (mock *mockedRepositoryRolSucces) AddPermiso(rolId, permisoId int64) error {
	return nil
}
func (mock *mockedRepositoryRolSucces) RemovePermiso(rolId, permisoId int64) error {
	return nil
}
func (mock *mockedRepositoryRolSucces) GetPermisos() (*[]domain.Permiso, error) {
	return nil, nil
}

//endregion

// region mock fail
type mockedRepositoryRolFail struct{}

func (mock *mockedRepositoryRolFail) CreateRol(domain.Rol) (insertedId int64, err error) {
	return 0, errors.New("Falla al crear el rol")
}
func (mock *mockedRepositoryRolFail) UpdateRol(domain.Rol) error {
	return errors.New("Falla al actualizar el rol")
}
func (mock *mockedRepositoryRolFail) GetRoles() (*[]domain.PermisosPorRol, error) {
	return nil, errors.New("Falla al obtener los roles")
}
func (mock *mockedRepositoryRolFail) GetRolByRolId(int64) (*domain.PermisosPorRol, error) {
	return nil, errors.New("Falla al obtener el rol")
}
func (mock *mockedRepositoryRolFail) DeleteRol(rolId int64) error {
	return errors.New("Falla al eliminar el rol")
}
func (mock *mockedRepositoryRolFail) ExistRol(name string) (bool, error) {
	return false, errors.New("Falla al verificar el rol")
}
func (mock *mockedRepositoryRolFail) AddPermiso(rolId, permisoId int64) error {
	return errors.New("Falla al crear el permiso")
}
func (mock *mockedRepositoryRolFail) RemovePermiso(rolId, permisoId int64) error {
	return errors.New("Falla al eliminar el permiso")
}
func (mock *mockedRepositoryRolFail) GetPermisos() (*[]domain.Permiso, error) {
	return nil, errors.New("Falla al obtener el permiso")
}

//endregion

// region mock empty
type mockedRepositoryRolEmpty struct{}

func (mock *mockedRepositoryRolEmpty) CreateRol(domain.Rol) (insertedId int64, err error) {
	return 1, nil
}
func (mock *mockedRepositoryRolEmpty) UpdateRol(domain.Rol) error {
	return nil
}
func (mock *mockedRepositoryRolEmpty) GetRoles() (*[]domain.PermisosPorRol, error) {
	return nil, nil
}
func (mock *mockedRepositoryRolEmpty) GetRolByRolId(int64) (*domain.PermisosPorRol, error) {
	return nil, nil
}
func (mock *mockedRepositoryRolEmpty) DeleteRol(rolId int64) error {
	return nil
}
func (mock *mockedRepositoryRolEmpty) ExistRol(name string) (bool, error) {
	return true, nil
}
func (mock *mockedRepositoryRolEmpty) AddPermiso(rolId, permisoId int64) error {
	return nil
}
func (mock *mockedRepositoryRolEmpty) RemovePermiso(rolId, permisoId int64) error {
	return nil
}
func (mock *mockedRepositoryRolEmpty) GetPermisos() (*[]domain.Permiso, error) {
	return nil, nil
}

//endregion

func TestCreateRol(t *testing.T) {
	//casos de prueba
	testCases := []testRol{
		{
			name: "Registros completos, sin errores",
			rol: domain.Rol{
				Nombre:      "administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoError,
				Datos: domain.Rol{
					Id:          1,
					Nombre:      "administrador",
					Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
				},
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, nombre menor a 6 caracteres",
			rol: domain.Rol{
				Nombre:      "aaa",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, descripcion menor a 50 caracteres",
			rol: domain.Rol{
				Nombre:      "Administrador",
				Descripcion: "Esta es una descripción de 49 caracteres,fallara",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, rol repetido", //mockear interfaz para devolver true y que crea que esta repetida
			rol: domain.Rol{
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.RegistroDuplicado,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolEmpty{},
		},
		{
			name: "Registros completos, error al crear el rol",
			rol: domain.Rol{
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &mockedRepositoryRolFail{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ErrorRepositorio,
				Datos:  nil,
			},
		},
		{
			name: "Registros incompletos, no se provee rol",
			rol: domain.Rol{
				Nombre:      "Administrador",
				Descripcion: "",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
		{
			name: "Registro incompletos, no se provee nombre",
			rol: domain.Rol{
				Nombre:      "",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
		{
			name: "Registro incompletos, no se provee nombre ni rol",
			rol: domain.Rol{
				Nombre:      "",
				Descripcion: "",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
	}

	//validador usado en main
	val := validacion.NewValidador()

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface, val)
			response := ser.CreateRol(tc.rol)

			//se valida si el error no es el esperado
			if response.Codigo != tc.expectedResult.Codigo {
				t.Errorf("expected %v, got %v", tc.expectedResult.Codigo, response.Codigo)
			}

			//se valida si los datos no son los esperados
			if response.Datos != tc.expectedResult.Datos {
				t.Errorf("expected %v, got %v", tc.expectedResult.Datos, response.Datos)
			}

		})
	}
}

func TestUpdateRol(t *testing.T) {
	testCases := []testRol{
		{
			name: "Registros completos, sin errores",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoError,
				Datos: &domain.Rol{
					Id:          1,
					Nombre:      "administrador",
					Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
				},
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, id menor a 1",
			rol: domain.Rol{
				Id:          0,
				Nombre:      "Administrador",
				Descripcion: "Esta es una descripción de 49 caracteres,fallara",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, nombre menor a 6 caracteres",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "aaa",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, descripcion menor a 50 caracteres",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "Administrador",
				Descripcion: "Esta es una descripción de 49 caracteres,fallara",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Registros completos, rol repetido", //mockear interfaz para devolver true y que crea que esta repetida
			rol: domain.Rol{
				Id:          1,
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.RegistroDuplicado,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolEmpty{},
		},
		{
			name: "Registros completos, error al actualizar el rol",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &mockedRepositoryRolFail{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ErrorRepositorio,
				Datos:  nil,
			},
		},
		{
			name: "Registros incompletos, no se provee rol",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "Administrador",
				Descripcion: "",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
		{
			name: "Registro incompletos, no se provee nombre",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
		{
			name: "Registro incompletos, no se provee id",
			rol: domain.Rol{
				Id:          0,
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
		{
			name: "Registro incompletos, no se provee nombre ni rol",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "",
				Descripcion: "",
			},
			mockedInterface: &mockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
	}

	//validador usado en main
	val := validacion.NewValidador()

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface, val)
			response := ser.UpdateRol(tc.rol)

			//se valida si el error no es el esperado
			if response.Codigo != tc.expectedResult.Codigo {
				t.Errorf("expected %v, got %v", tc.expectedResult.Codigo, response.Codigo)
			}

			if response.Datos != tc.expectedResult.Datos {
				t.Errorf("expected %v, got %v", tc.expectedResult.Datos, response.Datos)
			}
		})
	}
}

func TestGetRoles(t *testing.T) {
	testCases := []testRol{
		{
			name: "No errores en interfaces",
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoError,
				Datos:  &datosRoles,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Error en repositorio",
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ErrorRepositorio,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolFail{},
		},
		{
			name: "No hay datos disponibles",
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoData,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolEmpty{},
		},
	}

	//validador usado en main
	val := validacion.NewValidador()

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface, val)
			response := ser.GetRoles()

			//se valida si el error no es el esperado
			if response.Codigo != tc.expectedResult.Codigo {
				t.Errorf("expected %v, got %v", tc.expectedResult.Codigo, response.Codigo)
			}

			//se valida si los datos no son los esperados
			if response.Datos != tc.expectedResult.Datos {
				t.Errorf("expected %v, got %v", tc.expectedResult.Datos, response.Datos)
			}
		})
	}
}

func TestGetPermisosXRolByRolId(t *testing.T) {
	testCases := []testRol{
		{
			name: "No errores en interfaces",
			rol:  domain.Rol{Id: 1},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoError,
				Datos:  &datosRoles[0],
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Id formato incorrecto",
			rol:  domain.Rol{Id: 0}, //solo se usa el id
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolSucces{},
		},
		{
			name: "Error en repositorio",
			rol:  domain.Rol{Id: 1}, //solo se usa el id
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ErrorRepositorio,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolFail{},
		},
		{
			name: "No hay datos disponibles",
			rol:  domain.Rol{Id: 1}, //solo se usa el id
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoData,
				Datos:  nil,
			},
			mockedInterface: &mockedRepositoryRolEmpty{},
		},
	}

	//validador usado en main
	val := validacion.NewValidador()

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface, val)
			response := ser.GetPermisosXRolByRolId(tc.rol.Id)

			//se valida si el error no es el esperado
			if response.Codigo != tc.expectedResult.Codigo {
				t.Errorf("expected %v, got %v", tc.expectedResult.Codigo, response.Codigo)
			}

			//se valida si los datos no son los esperados
			if response.Datos != tc.expectedResult.Datos {
				t.Errorf("expected %v, got %v", tc.expectedResult.Datos, response.Datos)
			}
		})
	}
}

// func TestAddPermiso(t *testing.T) {
// 	testCases := []testRol{

// 	}
// }
