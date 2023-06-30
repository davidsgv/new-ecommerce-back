package domain

type Servidor struct {
	Id          uint   `validate:"min=0"`
	Dominio     string `validate:"min=3,max=500"`
	DireccionIp string `validate:"min=6,max=45"`
	NombreBD    string `validate:"min=3,max=100"`
	Usuario     string `validate:"min=5,max=100"`
	Password    string `validate:"min=15,max=100"`
	IdEmpresa   uint   `validate:"nonzero"`
}
