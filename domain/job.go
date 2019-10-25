package domain

import "log"

type Job struct {
	Datastore *Datastore
	Store     string
	CheckType string `yaml:"checkType"`
	Sql       string
	ToReturn  string `yaml:"toReturn"`
}


func (this *Job) Query() string {
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
			log.Fatal("Failed to scan row", err)
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
