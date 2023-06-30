package autenticacion

import "time"

type GetConexionesResponse struct {
	Id               uint      `json:"id"`
	Token            string    `json:"token"`
	Ingreso          time.Time `json:"ingreso"`
	Equipo           *string   `json:"equipo"`
	VencimientoToken time.Time `json:"vencimiento_token"`
	IdUsuario        uint      `json:"id_usuario"`
}
