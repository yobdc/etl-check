package domain

import (
	"database/sql"
	"log"
)

type Datastore struct {
	Uri           string
	Name          string
	DatastoreType string
	Db            *sql.DB
}

func (self *Datastore) Open() *sql.DB {
	db, err := sql.Open("mysql", self.Uri)
	if err != nil {
		log.Fatal("数据源打开错误%v", err)
	}
	self.Db = db
}


func (self *Datastore) Close() {
	self.Db.Close()
}
