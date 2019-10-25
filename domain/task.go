package domain

import (
	"log"
	"math"
	"strconv"
)

type Task struct {
	Name                 string
	Left                 *Job
	Right                *Job
	Delta                string
	Retry                string
	RetryInterval        string `yaml:"retryInterval"`
	RetryTimes           int    `yaml:"retryTimes"`
	RetryIntervalSeconds int    `yaml:"retryIntervalSeconds"`
	Op                   string
}

func (this *Task) Exec() (result bool) {
	var leftNum float64
	var rightNum float64
	leftResult := this.Left.Query()
	rightResult := this.Right.Query()

	switch this.Op {
	case "eq":
		result = leftResult == rightResult
	case "ne":
		result = leftResult != rightResult
	case "gt":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		result = leftNum > rightNum
	case "lt":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		result = leftNum < rightNum
	case "ge":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		result = leftNum >= rightNum
	case "le":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		result = leftNum <= rightNum
	case "gtin":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		deltaVal, _ := strconv.ParseFloat(this.Delta, 64)
		result = (leftNum - rightNum) < deltaVal
	case "ltin":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		deltaVal, _ := strconv.ParseFloat(this.Delta, 64)
		result = (rightNum - leftNum) < deltaVal
	case "in":
		leftNum, _ = strconv.ParseFloat(leftResult, 64)
		rightNum, _ = strconv.ParseFloat(rightResult, 64)
		deltaVal, _ := strconv.ParseFloat(this.Delta, 64)
		result = math.Abs(rightNum-leftNum) < deltaVal
	}
	log.Println("[Task]", this.Name, ": ", result)
	return result
}
