package main

import (
	"etl-check/domain"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	yamlFile := os.Args[1]
	//cmdArgs := os.Args[2]
	domain.Parse(yamlFile)

	//taskSet := tas
}
