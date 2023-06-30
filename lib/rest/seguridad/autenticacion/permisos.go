package autenticacion

type GetPermisosResponse struct {
	Id      uint   `json:"id"`
	Recurso string `json:"recurso"`
	Modulo  string `json:"modulo"`
}
