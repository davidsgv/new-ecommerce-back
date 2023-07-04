package mysql

import (
	"database/sql"
	"log"
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
