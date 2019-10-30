package domain

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

// TaskSet yaml对应工程
type TaskSet struct {
	Tasks      []Task
	Datastores []Datastore
	Vars       []EnvVar
	Variables  map[string]string
}

// AppReturn 定义整个app的返回值
var AppReturn string

// BuildEnvs 将环境变量放入Variables
func (conf *TaskSet) BuildEnvs() {
	conf.Variables = make(map[string]string)

	for i := 0; i < len(conf.Vars); i++ {
		varItem := &conf.Vars[i]
		varItem.Datastore.Open()
		conf.Variables[varItem.Name] = varItem.Query()
	}

	for i := range conf.Tasks {
		for k, v := range conf.Variables {
			conf.Tasks[i].Left.SQL = strings.ReplaceAll(conf.Tasks[i].Left.SQL, "{"+k+"}", v)
			conf.Tasks[i].Right.SQL = strings.ReplaceAll(conf.Tasks[i].Right.SQL, "{"+k+"}", v)
		}
	}
}

// Exec 执行工程中的所有任务
func (conf *TaskSet) Exec() {
	for i := 0; i < len(conf.Tasks); i++ {
		task := conf.Tasks[i]
		task.Exec()
	}
}

// Parse 解析配置文件，读取配置项
func Parse(configFile string) *TaskSet {
	conf := new(TaskSet)
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
	log.Println("conf", conf)

	dsMap := make(map[string]*Datastore)

	for i := 0; i < len(conf.Datastores); i++ {
		ds := conf.Datastores[i]
		dsMap[ds.Name] = &ds
	}

	for i := 0; i < len(conf.Vars); i++ {
		varItem := &conf.Vars[i]
		varItem.Datastore = dsMap[varItem.Store]
	}

	for i := 0; i < len(conf.Tasks); i++ {
		taskItem := &conf.Tasks[i]
		taskItem.Left.Datastore = dsMap[taskItem.Left.Store]
		taskItem.Right.Datastore = dsMap[taskItem.Right.Store]
	}

	return conf
}
