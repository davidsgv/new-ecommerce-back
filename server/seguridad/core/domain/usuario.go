package domain

type Usuario struct {
	Id        uint   `validate:"id"`
	Password  string `validate:"min=10,max=500"`
	Correo    string `validate:"max=500,email"`
	Telefono  string `validate:"number,max=10"`
	Celular   string `validate:"number,max=15"`
	Direccion string `validate:"max=200"`
	//Secreto   string
	Activo bool
}
