package domain

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type TaskSet struct {
	Tasks      []Task
	Datastores []Datastore
	Vars       []EnvVar
	Variables  map[string]string
}

func Parse(configFile string) *TaskSet {
	conf := new(TaskSet)
	yamlFile, err := ioutil.ReadFile(configFile)
	log.Println("yamlFile:", yamlFile)

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)

	dsMap := make(map[string]*Datastore)

	for i := 0; i < len(conf.Datastores); i++ {
		ds := conf.Datastores[i]
		dsMap[ds.Name] = &ds
	}

	for i := 0; i < len(conf.Vars); i++ {
		varItem := conf.Vars[i]
		varItem.Datastore = dsMap[varItem.Store]
	}

	for i := 0; i < len(conf.Tasks); i++ {
		taskItem := conf.Tasks[i]
		taskItem.Left.Datastore = dsMap[taskItem.Left.Store]
		taskItem.Right.Datastore = dsMap[taskItem.Right.Store]
	}
	return conf
}
