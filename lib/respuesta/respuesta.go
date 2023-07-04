// Este paquete se utiliza en todos los microservicios para entregar una respuesta estandarizada
package respuesta

// Estructura basica de la respuesta de un servicio
type Respuesta struct {
	Codigo Codigo
	Error  error
	Datos  interface{}
}

// Crea una nueva respuesta desde cero
// Solmente usar si las otras funciones no son suficientes para entregar la respuesta deseada
func NewRespuesta(codigo Codigo, err error, datos interface{}) *Respuesta {
	return &Respuesta{
		Codigo: codigo,
		Error:  err,
		Datos:  datos,
	}
}

// No hubo errores en el proceso
func NewNoErr(data interface{}) *Respuesta {
	return NewRespuesta(NoError, nil, data)
}

// no se encontraron datos
func NewNoData() *Respuesta {
	return NewRespuesta(NoData, nil, nil)
}

// Algún dato no paso la validación
func NewErrValidation(err error) *Respuesta {
	return NewRespuesta(ValidacionDatos, err, nil)
}

// Error en la implementación del repositorio
func NewErrRepository(err error) *Respuesta {
	return NewRespuesta(ErrorRepositorio, err, nil)
}

// Datos duplicados
func NewErrDuplicateRow(err error) *Respuesta {
	return NewRespuesta(RegistroDuplicado, err, nil)
}
