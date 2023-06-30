package mysql

import (
	"database/sql"
	"errors"
	"seguridad/core/domain"
)

func (repo *MysqlRepository) GetUsuarios() ([]domain.Usuario, error) {
	query := `
		SELECT id, correo, telefono, celular, direccion
		FROM usuario
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	usuarios := []domain.Usuario{}
	for rows.Next() {
		usuario := domain.Usuario{}

		rows.Scan(&usuario.Id, &usuario.Correo, &usuario.Telefono, &usuario.Celular, &usuario.Direccion)

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo *MysqlRepository) GetUsuariosById(usuarioId uint) ([]domain.Usuario, error) {
	query := `
		SELECT id, correo, telefono, celular, direccion
		FROM usuario
		WHERE id = ?
	`

	rows, err := repo.db.Query(query, usuarioId)

	if err != nil {
		return nil, err
	}

	usuarios := []domain.Usuario{}
	for rows.Next() {
		usuario := domain.Usuario{}

		rows.Scan(&usuario.Id, &usuario.Correo, &usuario.Telefono, &usuario.Celular, &usuario.Direccion)

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo *MysqlRepository) GetUsuariosByEmail(email string) (*domain.Usuario, error) {
	query := `
		SELECT id, correo, telefono, celular, direccion
		FROM usuario
		WHERE correo = ?
	`

	row := repo.db.QueryRow(query, email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	usuario := domain.Usuario{}

	err := row.Scan(&usuario.Id, &usuario.Correo, &usuario.Telefono, &usuario.Celular, &usuario.Direccion)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (repo *MysqlRepository) CreateUsuario(usuario domain.Usuario) (int64, error) {
	query := `
		INSERT INTO usuario (id, password, correo, telefono, celular, direccion)
		VALUES (?,?,?,?,?,?)
	`
	result, err := repo.db.Exec(query, usuario.Id, usuario.Password, usuario.Correo, usuario.Telefono, usuario.Celular, usuario.Direccion)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows <= 0 {
		return 0, errors.New("ninguna columna modificada")
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("error obteniendo el id")
	}
	return lastId, nil
}

func (repo *MysqlRepository) UpdateUsuario(usuario domain.Usuario) error {
	query := `
		UPDATE usuario
		SET correo = ?, telefono = ?, celular = ?, direccion = ?
		WHERE id = ?
	`
	result, err := repo.db.Exec(query, usuario.Correo, usuario.Telefono, usuario.Celular, usuario.Direccion, usuario.Id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows <= 0 {
		return errors.New("ninguna columna modificada")
	}
	return nil
}
