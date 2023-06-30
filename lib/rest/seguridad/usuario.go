package seguridad

type GetUsuariosResponse struct {
	Id        uint   `json:"id"`
	Correo    string `json:"correo"`
	Telefono  string `json:"telefono"`
	Celular   string `json:"celular"`
	Direccion string `json:"direccion"`
}

type CreateUsuariosRequest struct {
	Password  string `json:"password"`
	Correo    string `json:"correo"`
	Telefono  string `json:"telefono"`
	Celular   string `json:"celular"`
	Direccion string `json:"direccion"`
}

type UpdateUsuariosRequest struct {
	Id        uint   `json:"id"`
	Correo    string `json:"correo"`
	Telefono  string `json:"telefono"`
	Celular   string `json:"celular"`
	Direccion string `json:"direccion"`
}
