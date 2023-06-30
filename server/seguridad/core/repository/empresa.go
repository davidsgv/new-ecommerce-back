package repository

import "seguridad/core/domain"

type iServidor interface {
	CreateServidor(*domain.Servidor) (*domain.Servidor, error)
	DeleteServidor(uint) error
	GetServidores() ([]domain.Servidor, error)
	UpdateServidor(*domain.Servidor) (*domain.Servidor, error)
}

type iEmpresa interface {
	CreateEmpresa(domain.Empresa) error
	DeleteEmpresa(uint) error
	GetEmpresas() ([]domain.Empresa, error)
	GetEmpresaById(uint) (*domain.Empresa, error)
	UpdateEmpresa(domain.Empresa) error
}

type IRepositoryEmpresa interface {
	iServidor
	iEmpresa
}
