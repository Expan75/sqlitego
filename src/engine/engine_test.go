package engine

import (
	"testing"
)

func stringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func TestParseValidInsertStatement(t *testing.T) {
	result := parseInsertStatement("INSERT INTO tablename (col1, col2) VALUES (val1,val2)")
	
	if result.table != "tablename" {
		t.Fatalf("tablename parsed incorrectly")
	} else if result.values["col1"] != "val1" {
		t.Fatalf("columns and values parsed incorrectly")
	}
}