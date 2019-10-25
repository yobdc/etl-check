package domain

import (
	"fmt"
	"log"
)

type EnvVar struct {
	Name      string
	Store     string
	Datastore *Datastore
	Sql       string
}

func (this *EnvVar) Query() string {
	stmt, err := this.Datastore.Db.Prepare(this.Sql)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query()
	cols, err := rows.Columns()

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return "nil"
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

	}

	return result[0]
}
