package domain

import "log"

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

	if job.CheckType != "" {
		log.Printf("[job][%s] %s => %s\n", job.Store, job.SQL, result[0])
		if job.ToReturn != "" {
			AppReturn = job.ToReturn
		}
	}

	return result[0]
}
