package service_test

import (
	"respuesta"
	"seguridad/core/domain"
	"seguridad/core/repository"
	"seguridad/core/service"
	"testing"
)

// region test structs
type testRol struct {
	name            string
	rol             domain.Rol
	expectedResult  *respuesta.Respuesta
	mockedInterface repository.IRepositoryRol
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
				Datos: &domain.Rol{
					Id:          1,
					Nombre:      "administrador",
					Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
				},
			},
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolDuplicateError{},
		},
		{
			name: "Registros completos, error al crear el rol",
			rol: domain.Rol{
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &MockedRepositoryCreateRolError{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
	}

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface)
			response := ser.CreateRol(tc.rol)

			//se valida si el error no es el esperado
			if response.Codigo != tc.expectedResult.Codigo {
				t.Errorf("expected %v, got %v", tc.expectedResult.Codigo, response.Codigo)
			}

			//se valida si los datos no son los esperados
			if response.Datos == nil && tc.expectedResult.Datos == nil {
				return
			}

			expectedRol := response.Datos.(*domain.Rol)
			gotRol := tc.expectedResult.Datos.(*domain.Rol)
			if expectedRol.Id != gotRol.Id {
				t.Errorf("expected %v, got %v", expectedRol, gotRol)
			}

			if expectedRol.Nombre != gotRol.Nombre {
				t.Errorf("expected %v, got %v", expectedRol, gotRol)
			}

			if expectedRol.Descripcion != gotRol.Descripcion {
				t.Errorf("expected %v, got %v", expectedRol, gotRol)
			}

			if expectedRol.Permisos != nil {
				t.Errorf("expected %v, got %v", expectedRol, gotRol)
			}

			if gotRol.Permisos != nil {
				t.Errorf("expected %v, got %v", expectedRol, gotRol)
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolEmpty{},
		},
		{
			name: "Registros completos, error al actualizar el rol",
			rol: domain.Rol{
				Id:          1,
				Nombre:      "Administrador",
				Descripcion: "Este rol tiene privilegios administrativos, es decir cuenta con todos los permisos",
			},
			mockedInterface: &MockedRepositoryRolFail{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
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
			mockedInterface: &MockedRepositoryRolSucces{},
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ValidacionDatos,
				Datos:  nil,
			},
		},
	}

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface)
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
			mockedInterface: &MockedRepositoryRolSucces{},
		},
		{
			name: "Error en repositorio",
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.ErrorRepositorio,
				Datos:  nil,
			},
			mockedInterface: &MockedRepositoryRolFail{},
		},
		{
			name: "No hay datos disponibles",
			expectedResult: &respuesta.Respuesta{
				Codigo: respuesta.NoData,
				Datos:  nil,
			},
			mockedInterface: &MockedRepositoryRolEmpty{},
		},
	}

	//ejecucion casos de prueba
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			//se crea el servicio con el mock de la prueba
			ser := service.NewRolService(tc.mockedInterface)
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

// func TestGetPermisosXRolByRolId(t *testing.T) {
// 	testCases := []testRol{
// 		{
// 			name: "No errores en interfaces",
// 			rol:  domain.Rol{Id: 1},
// 			expectedResult: &respuesta.Respuesta{
// 				Codigo: respuesta.NoError,
// 				Datos:  &datosRoles[0],
// 			},
// 			mockedInterface: &MockedRepositoryRolSucces{},
// 		},
// 		{
// 			name: "Id formato incorrecto",
// 			rol:  domain.Rol{Id: 0}, //solo se usa el id
// 			expectedResult: &respuesta.Respuesta{
// 				Codigo: respuesta.ValidacionDatos,
// 				Datos:  nil,
// 			},
// 			mockedInterface: &MockedRepositoryRolSucces{},
// 		},
// 		{
// 			name: "Error en repositorio",
// 			rol:  domain.Rol{Id: 1}, //solo se usa el id
// 			expectedResult: &respuesta.Respuesta{
// 				Codigo: respuesta.ErrorRepositorio,
// 				Datos:  nil,
// 			},
// 			mockedInterface: &MockedRepositoryRolFail{},
// 		},
// 		{
// 			name: "No hay datos disponibles",
// 			rol:  domain.Rol{Id: 1}, //solo se usa el id
// 			expectedResult: &respuesta.Respuesta{
// 				Codigo: respuesta.NoData,
// 				Datos:  nil,
// 			},
// 			mockedInterface: &MockedRepositoryRolEmpty{},
// 		},
// 	}

// 	//ejecucion casos de prueba
// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			//se crea el servicio con el mock de la prueba
// 			ser := service.NewRolService(tc.mockedInterface)
// 			response := ser.GetPermisosXRolByRolId(tc.rol.Id)

// 			//se valida si el error no es el esperado
// 			if response.Codigo != tc.expectedResult.Codigo {
// 				t.Errorf("expected %v, got %v", tc.expectedResult.Codigo, response.Codigo)
// 			}

// 			//se valida si los datos no son los esperados
// 			if response.Datos != tc.expectedResult.Datos {
// 				t.Errorf("expected %v, got %v", tc.expectedResult.Datos, response.Datos)
// 			}
// 		})
// 	}
// }

// func TestAddPermiso(t *testing.T) {
// 	testCases := []testRol{

// 	}
// }
