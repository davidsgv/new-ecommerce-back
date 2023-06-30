package respuesta

type Respuesta struct {
	Codigo Codigo
	Error  error
	Datos  interface{}
}

func NewRespuesta(codigo Codigo, err error, datos interface{}) *Respuesta {
	return &Respuesta{
		Codigo: codigo,
		Error:  err,
		Datos:  datos,
	}
}

func NewNoErr(data interface{}) *Respuesta {
	return NewRespuesta(NoError, nil, data)
}

func NewNoData() *Respuesta {
	return NewRespuesta(NoData, nil, nil)
}

func NewErrValidation(err error) *Respuesta {
	return NewRespuesta(ValidacionDatos, err, nil)
}

func NewErrRepository(err error) *Respuesta {
	return NewRespuesta(ErrorRepositorio, err, nil)
}

func NewErrDuplicateRow(err error) *Respuesta {
	return NewRespuesta(RegistroDuplicado, err, nil)
}
