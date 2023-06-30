package repository

import "seguridad/core/domain"

type IRepositoryUsuario interface {
	GetUsuarios() ([]domain.Usuario, error)
	GetUsuariosById(uint) ([]domain.Usuario, error)
	GetUsuariosByEmail(string) (*domain.Usuario, error)
	CreateUsuario(domain.Usuario) (int64, error)
	UpdateUsuario(domain.Usuario) error
}
