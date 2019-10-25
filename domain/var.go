package domain

import "log"

type EnvVar struct {
	Name      string
	Store     string
	Datastore *Datastore
	Sql       string
}

func (self *EnvVar) Query() {
	stmt, err := self.Datastore.Db.Prepare(self.Sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query()
	//return rows.Next().
}
