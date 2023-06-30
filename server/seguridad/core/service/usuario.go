package service

import (

	//serRes "respuestaServicio"
	"seguridad/core/domain"
	"seguridad/core/repository"
	"validacion"
)

type UsuarioServicio struct {
	repo      repository.IRepositoryUsuario
	validador *validacion.Validador
}

func NewUsuarioServicio(repo repository.IRepositoryUsuario, validador *validacion.Validador) *UsuarioServicio {
	return &UsuarioServicio{
		repo:      repo,
		validador: validador,
	}
}

func (servicio *UsuarioServicio) GetUsuarios() ([]domain.Usuario, error) {
	usuarios, err := servicio.repo.GetUsuarios()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (servicio *UsuarioServicio) GetUsuariosById(id uint) ([]domain.Usuario, error) {
	err := servicio.validador.Verificar("id", id)
	if err != nil {
		return nil, err
	}
	usuarios, err := servicio.repo.GetUsuariosById(id)
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

// func (servicio *UsuarioServicio) CreateUsuario(usuario domain.Usuario) *serRes.Respuesta { //(*domain.Usuario, error) {
// 	//validar datos
// 	// err := servicio.validador.Verificar("optional,telefono", usuario.Celular)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	//revisar que no existe el registro
// 	usuarioExistente, err := servicio.repo.GetUsuariosByEmail(usuario.Correo)
// 	if err != nil {
// 		return serRes.NewRespuesta(serRes.ErrorInesperado, err, nil)
// 	}
// 	if usuarioExistente != nil {
// 		err = fmt.Errorf("El usuario con correo %s ya existe", usuario.Correo)
// 		return serRes.NewRespuesta(serRes.RegistroDuplicado, err, nil)
// 	}

// 	contraseña, err := cifrarContraseña(usuario.Password)
// 	if err != nil {
// 		return serRes.NewRespuesta(serRes.ErrorInesperado, err, nil)
// 	}
// 	usuario.Password = contraseña

// 	id, err := servicio.repo.CreateUsuario(usuario)
// 	if err != nil {
// 		return serRes.NewRespuesta(serRes.ErrorRepositorio, err, nil)
// 	}

// 	usuario.Id = uint(id)
// 	return serRes.NewRespuesta(serRes.NoError, nil, usuario)
// }

func (servicio *UsuarioServicio) UpdateUsuario(usuario domain.Usuario) (*domain.Usuario, error) {
	err := servicio.repo.UpdateUsuario(usuario)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
