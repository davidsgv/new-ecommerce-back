package respuesta

// Representa un error ocurrido en un servicio
type Codigo int

const (
	//no errores
	NoError             = 0
	RegistroCreado      = 1
	RegistroActualizado = 2
	RegistroBorrado     = 3
	NoData              = 4

	//error inesperado
	ErrorInesperado                    = 10
	RegistroCreadoErrorObteniendoDatos = 11

	//Validacion datos
	ValidacionDatos = 100
	BadBinding      = 101

	//error servicio
	RegistroDuplicado = 201

	//error base de datos (no deberian existir)
	ErrorRepositorio = 300
)
