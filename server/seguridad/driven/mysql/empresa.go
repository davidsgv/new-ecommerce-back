package mysql

// import (
// 	"errors"
// 	"log"
// 	"seguridad/core/domain"
// )

// func (repo *MysqlRepository) CreateServidor(servidor *domain.Servidor) (*domain.Servidor, error) {
// 	query := `
// 	INSERT INTO servidor (dominio, direccion_ip, nombre_bd, usuario, password, empresa_id)
// 	VALUES (?,?,?,?,?,?)
// 	`

// 	result, err := repo.db.Exec(query, servidor.Dominio, servidor.DireccionIp, servidor.NombreBD, servidor.Usuario, servidor.Password, servidor.IdEmpresa)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filasInsertadas, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if filasInsertadas <= 0 {
// 		return nil, errors.New("no se insertaron filas")
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}

// 	servidorCreado := domain.Servidor{
// 		Id:        uint(id),
// 		Dominio:   servidor.Dominio,
// 		IdEmpresa: servidor.IdEmpresa,
// 	}
// 	return &servidorCreado, nil
// }

// func (repo *MysqlRepository) UpdateServidor(servidor *domain.Servidor) (*domain.Servidor, error) {
// 	query := `
// 		UPDATE servidor
// 		SET dominio = ?, direccion_ip = ?, nombre_bd = ?, usuario = ?, password = ?, empresa_id = ?
// 		WHERE id = ?
// 	`

// 	result, err := repo.db.Exec(query, servidor.Dominio, servidor.DireccionIp, servidor.NombreBD,
// 		servidor.Usuario, servidor.Password, servidor.IdEmpresa, servidor.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filasInsertadas, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if filasInsertadas <= 0 {
// 		return nil, errors.New("no se insertaron filas")
// 	}

// 	servidorCreado := domain.Servidor{
// 		Id:        servidor.Id,
// 		Dominio:   servidor.Dominio,
// 		IdEmpresa: servidor.IdEmpresa,
// 	}
// 	return &servidorCreado, nil
// }

// func (repo *MysqlRepository) GetServidores() ([]domain.Servidor, error) {
// 	query := `
// 		SELECT id, dominio, direccion_ip, nombre_bd, usuario, password, empresa_id
// 		FROM servidor
// 	`

// 	rows, err := repo.db.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	servidores := []domain.Servidor{}
// 	for rows.Next() {
// 		servidor := domain.Servidor{}
// 		rows.Scan(&servidor.Id, &servidor.Dominio, &servidor.DireccionIp, &servidor.NombreBD, &servidor.Usuario, &servidor.Password, &servidor.IdEmpresa)
// 		servidores = append(servidores, servidor)
// 	}

// 	err = rows.Close()
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	return servidores, nil
// }

// // func (repo *MysqlRepository) GetEmpresas() ([]domain.Empresa, error){
// // 	query := `
// // 		SELECT
// // 			emp.id
// // 			, emp.logo
// // 			, emp.limite_usuario
// // 			, emp.razon_social
// // 			, emp.identificacion
// // 			, tid.id
// // 			, tid.nombre
// // 			, tid.codigo
// // 			, upe.activo
// // 			, usu.id
// // 			, usu. correo
// // 		FROM empresa emp
// // 		INNER JOIN tipo_identificacion tid
// // 			ON tid.id = emp.tipo_identificacion_id
// // 		INNER JOIN usuario_por_empresa upe
// // 			ON upe.empresa_id = emp.id
// // 		INNER JOIN usuario usu
// // 			ON usu.id = upe.usuario_id
// // 	`

// // 	rows, err := repo.db.Query(query)

// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	empresas := []domain.Empresa{}
// // 	for rows.Next() {
// // 		servidor := domain.Servidor{}
// // 		rows.Scan(&servidor.Id, &servidor.Dominio, &servidor.DireccionIp, &servidor.NombreBD, &servidor.Usuario, &servidor.Password, &servidor.IdEmpresa)
// // 		servidores = append(servidores, servidor)
// // 	}

// // 	err = rows.Close()
// // 	if err != nil {
// // 		log.Fatalln(err.Error())
// // 	}

// // 	return servidores, nil
// // }

// // func (repo *MysqlRepository) GetEmpresaById(uint) (*domain.Empresa, error){

// // }
