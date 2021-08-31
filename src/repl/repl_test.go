package repl

import (
	"testing"
)


func TestParseInput(t *testing.T) {
	input := "input    "
	output := parseInput("input    ")
	if len(output) != 5 {
        t.Fatalf("parseInput yielded non formatted output, parseInput(%s)", input)
    }
}
