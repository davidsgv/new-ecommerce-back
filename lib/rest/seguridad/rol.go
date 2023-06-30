package seguridad

type CreateRolRequest struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type CreateRolResponse struct {
	Id          int64  `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

// type GetRolByRolIdRequest struct {
// 	RolId int64
// }

type AuxGetRolesResponse struct {
	Id        int64  `json:"id"`
	Modulo    string `json:"modulo"`
	Recurso   string `json:"recurso"`
	Editar    bool   `json:"editar"`
	Consultar bool   `json:"consultar"`
	Crear     bool   `json:"crear"`
	Borrar    bool   `json:"borrar"`
}

type GetRolesResponse struct {
	Id          int64                 `json:"id"`
	Nombre      string                `json:"nombre"`
	Descripcion string                `json:"descripcion"`
	Permisos    []AuxGetRolesResponse `json:"permisos"`
}
