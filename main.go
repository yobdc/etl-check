package main

import (
	"fmt"
	"github.com/yobdc/etl-check/domain"
	"os"
)

func main() {
	fmt.Println(os.Args)
	yamlFile := os.Args[1]
	//cmdArgs := os.Args[2]
	taskSet := domain.Parse(yamlFile)
	taskSet.BuildEnvs()
	taskSet.Exec()
	//taskSet := tas
}
