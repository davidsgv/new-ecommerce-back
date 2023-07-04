package mysql

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"seguridad/core/domain"
// 	"strings"
// 	"time"

// 	"golang.org/x/exp/slices"
// )

// func (repo *MysqlRepository) GetUsuarioLogin(correo, dominio string) (*domain.Usuario, error) {
// 	var query string = `
// 		SELECT usu.id, correo, usu.password, activo
// 		FROM usuario usu
// 		INNER JOIN usuario_por_empresa upe
// 			ON usu.id = upe.usuario_id
// 		INNER JOIN empresa emp
// 			ON emp.id = upe.empresa_id
// 		INNER JOIN servidor ser
// 			ON emp.id = ser.empresa_id
// 		WHERE usu.correo = ?
// 			AND upe.activo = 1
// 			AND ser.dominio = ?
// 	`

// 	usuario := domain.Usuario{}
// 	row := repo.db.QueryRow(query, correo, dominio)
// 	err := row.Scan(&usuario.Id, &usuario.Correo, &usuario.Password, &usuario.Activo)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return nil, nil
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	return &usuario, nil
// }

// func (repo *MysqlRepository) SaveConexion(conexion *domain.Conexion) error {
// 	query := `
// 		INSERT INTO conexion (token, ingreso, equipo, vencimiento_token, usuario_id)
// 		VALUES (?,?,?,?,?)
// 	`
// 	result, err := repo.db.Exec(query, conexion.Token, conexion.Ingreso, conexion.Equipo, conexion.VencimientoToken, conexion.IdUsuario)
// 	if err != nil {
// 		return err
// 	}
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rows <= 0 {
// 		return errors.New("ninguna columna modificada")
// 	}
// 	return nil
// }

// func (repo *MysqlRepository) GetRolesByUsuarioId(id uint) ([]domain.Rol, error) {
// 	query := `
// 		SELECT rol.id, rol.nombre, rol.descripcion
// 		FROM rol
// 		INNER JOIN usuario_por_rol upr
// 			ON upr.rol_id = rol.id
// 		WHERE usuario_id = ?
// 	`

// 	rows, err := repo.db.Query(query, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	roles := []domain.Rol{}
// 	for rows.Next() {
// 		rol := domain.Rol{}
// 		rows.Scan(&rol.Id, &rol.Nombre, &rol.Descripcion)
// 		roles = append(roles, rol)
// 	}

// 	return roles, nil
// }

// func (repo *MysqlRepository) ValidarUsuario(correo, dominio string, fechaFirma time.Time) (bool, error) {
// 	var usuarioId int
// 	query := `
// 		SELECT usu.id
// 		FROM usuario usu
// 		INNER JOIN usuario_por_empresa upe
// 			ON upe.usuario_id = usu.id
// 		INNER JOIN empresa emp
// 			ON emp.id = upe.empresa_id
// 		INNER JOIN servidor ser
// 			ON ser.empresa_id = emp.id
// 		WHERE usu.correo = ?
// 			AND upe.activo = 1
// 			AND ser.dominio = ?
// 	`

// 	row := repo.db.QueryRow(query, correo, dominio)
// 	err := row.Scan(&usuarioId)
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return false, nil
// 	}
// 	if err != nil {
// 		return false, err
// 	}

// 	var hora sql.NullTime
// 	query = `
// 		SELECT hora
// 		FROM bloqueo blo
// 		INNER JOIN usuario
// 			ON blo.usuario_id = usuario.id
// 		WHERE ? <= hora
// 	`
// 	row = repo.db.QueryRow(query, fechaFirma)
// 	err = row.Scan(&hora)

// 	//si no hay registros no hay bloqueos del usuario
// 	if errors.Is(err, sql.ErrNoRows) {
// 		return true, nil
// 	}
// 	if err != nil {
// 		return false, err
// 	}

// 	return false, nil
// }

// func (repo *MysqlRepository) ValidarPermisos(recurso, modulo, operacion string, roles []string) (bool, error) {
// 	//se actualiza la operacion a la columna de la BD
// 	switch strings.ToLower(operacion) {
// 	case "post":
// 		operacion = "crear"
// 	case "get":
// 		operacion = "consultar"
// 	case "put":
// 		operacion = "editar"
// 	case "delete":
// 		operacion = "eliminar"
// 	}

// 	query := fmt.Sprintf(`
// 		SELECT rol.nombre
// 		FROM permiso per
// 		INNER JOIN permiso_por_rol ppr
// 			ON per.id = ppr.permiso_id
// 		INNER JOIN rol
// 			ON rol.id = ppr.rol_id
// 		WHERE per.modulo = ?
// 			AND per.recurso = ?
// 			AND ppr.%s = 1
// 	`, operacion)

// 	rows, err := repo.db.Query(query, modulo, recurso)
// 	if err != nil {
// 		return false, err
// 	}

// 	var rolEncontrado bool
// 	for rows.Next() {
// 		permiso := domain.Rol{}

// 		rows.Scan(&permiso.Nombre)

// 		//si no tiene el rol
// 		if slices.Index(roles, permiso.Nombre) == -1 {
// 			continue
// 		}

// 		rolEncontrado = true
// 	}

// 	if rolEncontrado {
// 		return true, nil
// 	}
// 	return false, nil
// }

// // func (repo *MysqlRepository) GetRoles() ([]domain.PermisosPorRol, error) {
// // 	//se actualiza la operacion a la columna de la BD
// // 	query := `
// // 		SELECT rol.id, rol.nombre, rol.descripcion, per.id, per.modulo, per.recurso, ppr.crear, ppr.editar, ppr.eliminar, ppr.consultar
// // 		FROM rol
// // 		INNER JOIN permiso_por_rol ppr
// // 			ON rol.id = ppr.rol_id
// // 		INNER JOIN permiso per
// // 			ON ppr.permiso_id = per.id`

// // 	rows, err := repo.db.Query(query)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	roles := []domain.PermisosPorRol{}
// // 	for rows.Next() {
// // 		rol := domain.PermisosPorRol{
// // 			Permisos: []domain.Permiso{},
// // 			Rol:      domain.Rol{},
// // 		}

// // 		rows.Scan(&rol.Rol.Id, &rol.Rol.Nombre, &rol.Rol.Descripcion, &rol.Permiso.Id, &rol.Permiso.Modulo, &rol.Permiso.Recurso, &rol.Crear, &rol.Editar, &rol.Eliminar, &rol.Consultar)

// // 		roles = append(roles, rol)
// // 	}

// // 	return roles, nil
// // }

// // func (repo *MysqlRepository) GetRolById(id uint) ([]domain.PermisoPorRol, error) {
// // 	//se actualiza la operacion a la columna de la BD
// // 	query := `
// // 		SELECT rol.id, rol.nombre, rol.descripcion, per.id, per.modulo, per.recurso, ppr.crear, ppr.editar, ppr.eliminar, ppr.consultar
// // 		FROM rol
// // 		INNER JOIN permiso_por_rol ppr
// // 			ON rol.id = ppr.rol_id
// // 		INNER JOIN permiso per
// // 			ON ppr.permiso_id = per.id
// // 		WHERE rol.id = ?
// // 	`

// // 	rows, err := repo.db.Query(query, id)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	roles := []domain.PermisoPorRol{}
// // 	for rows.Next() {
// // 		rol := domain.PermisoPorRol{
// // 			Permiso: domain.Permiso{},
// // 			Rol:     domain.Rol{},
// // 		}

// // 		rows.Scan(&rol.Rol.Id, &rol.Rol.Nombre, &rol.Rol.Descripcion, &rol.Permiso.Id, &rol.Permiso.Modulo, &rol.Permiso.Recurso, &rol.Crear, &rol.Editar, &rol.Eliminar, &rol.Consultar)

// // 		roles = append(roles, rol)
// // 	}

// // 	return roles, nil
// // }

// func (repo *MysqlRepository) BloquearSesion(correo string) error {
// 	query := `
// 		INSERT INTO bloqueo (hora, usuario_id)
// 		VALUES (
// 			NOW(),
// 			(SELECT id FROM usuario WHERE correo = ?)
// 		)
// 	`

// 	result, err := repo.db.Exec(query, correo)
// 	if err != nil {
// 		return err
// 	}
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rows <= 0 {
// 		return errors.New("no se insertaron registros")
// 	}
// 	return nil
// }

// func (repo *MysqlRepository) GetConexiones() ([]domain.Conexion, error) {
// 	query := `
// 		SELECT id, token, ingreso, equipo, vencimiento_token, usuario_id
// 		FROM conexion
// 	`

// 	rows, err := repo.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conexiones := []domain.Conexion{}
// 	for rows.Next() {
// 		conexion := domain.Conexion{}

// 		rows.Scan(&conexion.Id, &conexion.Token, &conexion.Ingreso, &conexion.Equipo, &conexion.VencimientoToken, &conexion.IdUsuario)

// 		conexiones = append(conexiones, conexion)
// 	}

// 	return conexiones, nil
// }

// func (repo *MysqlRepository) GetConexionesByUsuarioId(usuarioId uint) ([]domain.Conexion, error) {
// 	query := `
// 		SELECT id, token, ingreso, equipo, vencimiento_token, usuario_id
// 		FROM conexion
// 		WHERE usuario_id = ?
// 	`

// 	rows, err := repo.db.Query(query, usuarioId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conexiones := []domain.Conexion{}
// 	for rows.Next() {
// 		conexion := domain.Conexion{}

// 		rows.Scan(&conexion.Id, &conexion.Token, &conexion.Ingreso, &conexion.Equipo, &conexion.VencimientoToken, &conexion.IdUsuario)

// 		conexiones = append(conexiones, conexion)
// 	}

// 	return conexiones, nil
// }
