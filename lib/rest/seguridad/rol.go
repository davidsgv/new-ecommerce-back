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

type GetRolesResponsePermisos struct {
	Crear     bool   `json:"crear"`
	Editar    bool   `json:"editar"`
	Consultar bool   `json:"consultar"`
	Eliminar  bool   `json:"eliminar"`
	Id        int64  `json:"id"`
	Modulo    string `json:"modulo"`
	Recurso   string `json:"recurso"`
}

type GetRolesResponse struct {
	Id          int64                      `json:"id"`
	Nombre      string                     `json:"nombre"`
	Descripcion string                     `json:"descripcion"`
	Permisos    []GetRolesResponsePermisos `json:"permisos"`
}
