package autenticacion

type Permiso struct {
	Id        uint
	Modulo    string
	Recurso   string
	Crear     bool
	Eliminar  bool
	Consultar bool
	Editar    bool
}

type RolByIdResponse struct {
	Id          uint   `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Permisos    []Permiso
}
