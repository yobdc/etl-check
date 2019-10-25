package domain

type Task struct {
	Name  string
	Left  *Job
	Right *Job
}
