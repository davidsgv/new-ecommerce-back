package mysql

import (
	"database/sql"
	"errors"
	"log"
	"seguridad/core/domain"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	ConnMaxLifetime time.Duration //SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	ConnMaxIdleTime time.Duration //SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	MaxIdleConns    int           //SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	MaxOpenConns    int           //SetMaxOpenConns sets the maximum number of open connections to the database.
	URL             string        //connection string BD
}

type MysqlRepository struct {
	db *sql.DB
}

// func NewMysqlRepository(conf MysqlConfig) (repository.IRepositoryEmpresa, error) {
func NewMysqlRepository(conf MysqlConfig) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", conf.URL)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Second * conf.ConnMaxLifetime)
	db.SetConnMaxIdleTime(time.Second * conf.ConnMaxIdleTime)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("ERROR: NO BD")
		return nil, pingErr
	}

	return &MysqlRepository{
		db: db,
	}, nil
}

func (repo *MysqlRepository) CreateEmpresa(empresa domain.Empresa) error {
	return errors.New("NOT IMPLEMENTED")
}

// Empresa
func (repo *MysqlRepository) GetEmpresas() (empresas []domain.Empresa, err error) {
	var query string = `
		SELECT
			empresa.id,
			empresa.logo,
			empresa.limite_usuarios,
			empresa.nit,
			empresa.razon_social,
			tipo.codigo
		FROM empresa
		INNER JOIN tipo_identificacion tipo
			ON empresa.tipo_identificacion_id = tipo.id
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		empresa := domain.Empresa{}
		rows.Scan(&empresa.Id, &empresa.Logo, &empresa.LimiteUsuarios, &empresa.Nit,
			&empresa.RazonSocial, &empresa.Identificacion)
		empresas = append(empresas, empresa)
	}

	return empresas, nil
}

func (repo *MysqlRepository) GetEmpresaById(id uint) (*domain.Empresa, error) {
	var query string = `
		SELECT
			empresa.id,
			empresa.logo,
			empresa.limite_usuarios,
			empresa.nit,
			empresa.razon_social,
		FROM empresa
		WHERE empresa.id = ?
	`
	rows := repo.db.QueryRow(query, id)
	empresa := domain.Empresa{}
	err := rows.Scan(&empresa.Id, &empresa.Logo, &empresa.LimiteUsuarios, &empresa.Nit,
		&empresa.RazonSocial)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &empresa, nil
}

func (repo *MysqlRepository) UpdateEmpresa(empresa domain.Empresa) error {
	return errors.New("NOT IMPLEMENTED")
}

func (repo *MysqlRepository) DeleteEmpresa(id uint) error {
	return errors.New("NOT IMPLEMENTED")
}

func (repo *MysqlRepository) DeleteServidor(id uint) error {
	return errors.New("NOT IMPLEMENTED")
}
