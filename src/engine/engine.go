package engine

import (
	"fmt"
)

type selectParameters struct {
	table string
	columns []string
	conditions []string
}


func HandleSelectStatement(selectStatement string) {
    fmt.Println("A select statement was fed in!")
}

func parseSelectStatement(rawSelect string) selectParameters { return selectParameters() }
func validateSelectParameters(parameters selectParameters) selectParameters { return parameters }
func executeSelectStatement(validatedParameters selectParameters) {}

func HandleInsertStatement(selectStatement string) {
    fmt.Println("An insert statement was fed in!")
}