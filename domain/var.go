package domain

import (
	"fmt"
	"log"
)

// EnvVar 环境变量
type EnvVar struct {
	Name      string
	Store     string
	Datastore *Datastore
	SQL       string `yaml:"sql"`
}

// Query 执行sql查询，读取返回结果
func (envVar *EnvVar) Query() string {
	stmt, err := envVar.Datastore.Db.Prepare(envVar.SQL)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query()
	cols, err := rows.Columns()

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row,", err)
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
