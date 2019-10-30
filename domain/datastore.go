package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //go-mysql
	"log"
)

// Datastore 数据源
type Datastore struct {
	URI    string `yaml:"uri"`
	Name   string
	DbType string `yaml:"dbType"`
	Db     *sql.DB
}

// Open 打开数据库连接
func (datastore *Datastore) Open() *sql.DB {
	db, err := sql.Open(datastore.DbType, datastore.URI)
	if err != nil {
		log.Fatal("数据源打开错误", err)
	}
	datastore.Db = db
	return db
}

// Close 关闭数据库连接
func (datastore *Datastore) Close() {
	datastore.Db.Close()
}
