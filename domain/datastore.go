package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Datastore struct {
	Uri    string
	Name   string
	DbType string `yaml:"dbType"`
	Db     *sql.DB
}

func (this *Datastore) Open() *sql.DB {
	db, err := sql.Open(this.DbType, this.Uri)
	if err != nil {
		log.Fatal("数据源打开错误%v", err)
	}
	this.Db = db
	return db
}

func (this *Datastore) Close() {
	this.Db.Close()
}
