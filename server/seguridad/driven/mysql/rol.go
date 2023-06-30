package mysql

import (
	"errors"
	"seguridad/core/domain"
)

func (repo *MysqlRepository) CreateRol(rol domain.Rol) (insertedId int64, err error) {
	query := `
		INSERT INTO rol (nombre, descripcion)
		VALUES (?,?)
	`
	result, err := repo.db.Exec(query, rol.Nombre, rol.Descripcion)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows <= 0 {
		return 0, errors.New("Ninguna columna modificada")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("No fue posible obtener el id")
	}

	return id, nil
}

func (repo *MysqlRepository) UpdateRol(rol domain.Rol) error {
	query := `
		UPDATE rol
		SET nombre = ?, descripcion = ?
		WHERE id = ?
	`
	result, err := repo.db.Exec(query, rol.Nombre, rol.Descripcion, rol.Id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows <= 0 {
		return errors.New("Ninguna columna modificada")
	}

	return nil
}

func (repo *MysqlRepository) GetRoles() (*[]domain.PermisosPorRol, error) {
	query := `
		SELECT 
			rol.id
			, rol.nombre
			, rol.descripcion
			, ppr.crear
			, ppr.editar
			, ppr.eliminar
			, ppr.consultar
			, per.id
			, per.modulo
			, per.recurso
		FROM rol
		INNER JOIN permiso_por_rol ppr
			ON ppr.rol_id = rol.id
		INNER JOIN permiso per
			ON per.id = ppr.permiso_id
	`

	rows, err := repo.db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	permisosRoles := []domain.PermisosPorRol{}
	for rows.Next() {
		pr := domain.PermisosPorRol{}

		rows.Scan(&pr.Id, &pr.Nombre, &pr.Descripcion,
			&pr.Crear, &pr.Editar, &pr.Eliminar, &pr.Eliminar,
			pr.Permiso.Id, &pr.Permiso.Modulo, &pr.Permiso.Recurso)

		permisosRoles = append(permisosRoles, pr)
	}

	return &permisosRoles, nil
}

func (repo *MysqlRepository) GetRolByRolId(rolId int64) (*[]domain.PermisosPorRol, error) {
	query := `
		SELECT 
			rol.id
			, rol.nombre
			, rol.descripcion
			, ppr.crear
			, ppr.editar
			, ppr.eliminar
			, ppr.consultar
			, per.id
			, per.modulo
			, per.recurso
		FROM rol
		INNER JOIN permiso_por_rol ppr
			ON ppr.rol_id = rol.id
		INNER JOIN permiso per
			ON per.id = ppr.permiso_id
		WHERE rol.id = ?
	`

	rows, err := repo.db.Query(query, rolId)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	permisosRoles := []domain.PermisosPorRol{}
	for rows.Next() {
		pr := domain.PermisosPorRol{}

		rows.Scan(&pr.Id, &pr.Nombre, &pr.Descripcion,
			&pr.Crear, &pr.Editar, &pr.Eliminar, &pr.Eliminar,
			pr.Permiso.Id, &pr.Permiso.Modulo, &pr.Permiso.Recurso)

		permisosRoles = append(permisosRoles, pr)
	}

	return &permisosRoles, nil
}

func (repo *MysqlRepository) DeleteRol(rolId int64) error {
	return errors.New("Not implemented yet")
}

func (repo *MysqlRepository) AddPermiso(rolId, permisoId int64) error {
	return errors.New("Not implemented yet")
}
func (repo *MysqlRepository) RemovePermiso(rolId, permisoId int64) error {
	return errors.New("Not implemented yet")
}
func (repo *MysqlRepository) ExistRol(name string) (bool, error) {
	return false, errors.New("Not implemented yet")
}
func (repo *MysqlRepository) GetPermisos() (*[]domain.Permiso, error) {
	query := `
		SELECT id, modulo, recurso
		FROM permiso
	`

	rows, err := repo.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	permisos := []domain.Permiso{}
	for rows.Next() {
		permiso := domain.Permiso{}

		rows.Scan(&permiso.Id, &permiso.Modulo, &permiso.Recurso)

		permisos = append(permisos, permiso)
	}

	return &permisos, nil
}
