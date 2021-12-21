package engine

import (
	"fmt"
	"strings"
	"regexp"
)

type row struct {
	id uint64
	name string
	email string
}

var table = []row{}

type insertParameters struct {
	table string
	values map[string]string
}

func parseInsertStatement(rawInsertStatement string) insertParameters {
	params := insertParameters{
		strings.Split(rawInsertStatement, " ")[2],
		make(map[string]string),	
	}
	trimmedStatment := strings.Replace(rawInsertStatement, " ", "", -1)

	re := regexp.MustCompile(`\(([^\)]+)\)`)
	matches := re.FindAll([]byte(trimmedStatment), -1)
	columns := strings.TrimSuffix(string(matches[0])[1:], ")")
	values := strings.TrimSuffix(string(matches[1])[1:], ")")

	for i, column := range strings.Split(columns, ",") {
		params.values[column] = strings.Split(values, ",")[i]
	}

	return params
}

//func validateInsertParameters(parameters insertParameters) insertParameters { return insertParameters{} }
//func executeInsertStatement(validatedParameters insertParameters) {}

func HandleInsertStatement(insertStatement string) {
    fmt.Println("An insert statement was fed in!")
	parsedParams := parseInsertStatement(insertStatement)
	fmt.Println(parsedParams)
}

func HandleSelectStatement(selectStatement string) {
    fmt.Println("A select statement was fed in!")
}

//func parseSelectStatement(rawSelect string) selectParameters { return selectParameters{} }
//func validateSelectParameters(parameters selectParameters) selectParameters { return selectParameters{} }
//func executeSelectStatement(validatedParameters selectParameters) {}