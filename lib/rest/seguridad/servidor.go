package seguridad

// GET
type GetServidorResponse struct {
	Id        uint   `json:"id"`
	Dominio   string `json:"dominio"`
	EmpresaId uint   `json:"empresa_id"`
}

// POST
type CreateServidorRequest struct {
	Dominio     string `json:"dominio" binding:"min=3,max=500"`
	DireccionIp string `json:"direccion_ip" binding:"min=6,max=45"`
	NombreBD    string `json:"nombre_bd" binding:"min=3,max=100"`
	Usuario     string `json:"usuario" binding:"min=5,max=100"`
	Password    string `json:"password" binding:"min=15,max=100"`
	IdEmpresa   uint   `json:"id_empresa" binding:"required"`
}

type CreateServidorResponse struct {
	Id        uint   `json:"id"`
	Dominio   string `json:"dominio"`
	IdEmpresa uint   `json:"id_empresa"`
}

// PUT
type UpdateServidorRequest struct {
	Id uint `json:"id" binding:"required"`
	CreateServidorRequest
}

type UpdateServidorResponse struct {
	CreateServidorResponse
}
