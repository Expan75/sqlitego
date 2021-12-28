package engine

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type row struct {
	id    uint64
	name  string
	email string
}

var table = []row{}

type insertParameters struct {
	table        string
	columnValues map[string]string
}

func cleanString(text string) string {
	text = strings.ReplaceAll(text, "(", "")
	text = strings.ReplaceAll(text, ")", "")
	return strings.TrimSpace(text)
}

func parseInsert(rawInsertStatement string) (insertParameters, error) {
	insertParams := insertParameters{
		strings.Split(rawInsertStatement, " ")[2],
		make(map[string]string),
	}
	// parse the columns and their assigned values
	re := regexp.MustCompile(`\(.*?\)`)
	matches := re.FindAll([]byte(rawInsertStatement), -1)
	if len(matches) == 2 {
		cleanColumnsString := cleanString(string(matches[0]))
		cleanValuesString := cleanString(string(matches[1]))
		columns := strings.Split(cleanColumnsString, ",")
		values := strings.Split(cleanValuesString, ",")
		if len(columns) == len(values) {
			for i := 0; i < len(values); i++ {
				insertParams.columnValues[columns[i]] = values[i]
			}
		}
	}

	if len(insertParams.columnValues) == 0 {
		return insertParams, errors.New("invalid insert statement")
	} else {
		return insertParams, nil
	}
}

func validateInsert(parsedInsert insertParameters) (insertParameters, error) {
	if parsedInsert.table != "users" {
		return parsedInsert, errors.New(fmt.Sprintf("no table named :: %s", parsedInsert.table))
	}
	return parsedInsert, nil
}

func executeInsert(parsedInsert insertParameters) (string, error) {
	return "hej", errors.New("hej")
}

//func validateInsertParameters(parameters insertParameters) insertParameters { return insertParameters{} }
//func executeInsertStatement(validatedParameters insertParameters) {}

func HandleInsertStatement(insertStatement string) {
	parsedInsert, parseError := parseInsert(insertStatement)
	if parseError != nil {
		return
	}
	fmt.Println("parsing of insert yielded: ")
	fmt.Println(parsedInsert)
}

func HandleSelectStatement(selectStatement string) {
	fmt.Println("A select statement was fed in!")
}

//func parseSelectStatement(rawSelect string) selectParameters { return selectParameters{} }
//func validateSelectParameters(parameters selectParameters) selectParameters { return selectParameters{} }
//func executeSelectStatement(validatedParameters selectParameters) {}
