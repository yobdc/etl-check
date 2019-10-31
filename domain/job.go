package domain

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

const (
	// NullString null常量
	NullString = "(null)"
)

// Job 最小sql执行工作
type Job struct {
	Datastore *Datastore
	Store     string
	CheckType string `yaml:"checkType"`
	SQL       string `yaml:"sql"`
	ToReturn  string `yaml:"toReturn"`
}

// Query 执行sql查询语句
func (job *Job) Query() string {
	stmt, err := job.Datastore.Db.Prepare(job.SQL)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(strings.ToLower(job.SQL), "select") {
		rows, err := stmt.Query()
		cols, err := rows.Columns()
		vals := make([]interface{}, len(cols))
		for i := range cols {
			vals[i] = new(sql.RawBytes)
		}

		for rows.Next() {
			err = rows.Scan(vals...)
			if err != nil {
				log.Fatal("Failed to scan row", err)
				// panic(err)
				return "nil"
			}

			if job.CheckType == "" {
				log.Printf("[job][%s] %s => %s\n", job.Store, job.SQL, vals[0])
				if job.ToReturn != "" {
					AppReturn = job.ToReturn
				}
			} else if job.CheckType == "fieldType" {
			} else if job.CheckType == "mappedFieldType" {
			}
		}
		return fmt.Sprintf("%s", vals[0])
	}
	execResult, err := stmt.Exec()
	if err != nil {
		log.Fatal("Exec failed,err:", err)
	}
	rowsCount, err := execResult.RowsAffected()
	log.Printf("[job][%s] %s => updated %d rows", job.Store, job.SQL, rowsCount)
	return NullString
}
