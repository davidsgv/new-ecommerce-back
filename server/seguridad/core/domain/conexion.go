package domain

import "time"

type Conexion struct {
	Id               uint
	Token            string
	Ingreso          time.Time
	Equipo           *string
	VencimientoToken time.Time
	IdUsuario        uint
}
