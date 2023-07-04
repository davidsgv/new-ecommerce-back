package domain

import (
	"fmt"
	"validator"
)

type Rol struct {
	Id          int64
	Nombre      string
	Descripcion string
	Permisos    []PermisosPorRol
}

func (rol *Rol) ValidateId() error {
	err := validator.Validate("id", rol.Id)
	if err != nil {
		return fmt.Errorf("Id: %w", err)
	}
	return nil
}

func (rol *Rol) ValidateNombre() error {
	err := validator.Validate("min=6,max=200", rol.Nombre)
	if err != nil {
		return fmt.Errorf("Nombre: %w", err)
	}
	return nil
}

func (rol *Rol) ValidateDescripcion() error {
	err := validator.Validate("min=50,max=2000", rol.Descripcion)
	if err != nil {
		return fmt.Errorf("Descripción: %w", err)
	}
	return nil
}
