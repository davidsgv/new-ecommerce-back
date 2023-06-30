package domain

type Empresa struct {
	Id             uint
	Identificacion string
	Servidores     []Servidor
	Logo           string `validate:"max:1000"`
	LimiteUsuarios int    `validate:"nonzero"`
	Nit            string `validate:"nonzero,max=45"`
	RazonSocial    string `validate:"nonzero,max=500"`
	Usuarios       []Usuario
}
