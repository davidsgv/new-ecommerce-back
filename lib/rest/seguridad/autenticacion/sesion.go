package autenticacion

type SesionRequest struct {
	Token     string `json:"token" binding:"required"`
	Modulo    string `json:"modulo" binding:"required"`
	Recurso   string `json:"recurso" binding:"required"`
	Operacion string `json:"operacion" binding:"required"`
}

type SesionResponse struct {
	Permiso bool `json:"permiso"`
}
