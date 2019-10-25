package domain

type Job struct {
	Datastore *Datastore
	Store     string
	CheckType string
	Sql       string
	ToReturn  string
}

func (job *Job) Query() {

}
