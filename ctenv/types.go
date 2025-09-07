package ctenv

import (
	"database/sql"
	"gorm.io/gorm"
)

type DBType string

const (
	MySQL    DBType = "mysql"
	MariaDB  DBType = "mariadb"
	Postgres DBType = "postgres"
)

type DBContainer struct {
	container   Container
	cfg         Config
	sqlDB       *sql.DB
	gormDB      *gorm.DB
	exposedPort string
	driverName  string
}

func (d *DBContainer) DB() *sql.DB { return d.sqlDB }
func (d *DBContainer) GormDB()     { return d.gormDB }
