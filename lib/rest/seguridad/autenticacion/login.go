package autenticacion

type LoginRequest struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
