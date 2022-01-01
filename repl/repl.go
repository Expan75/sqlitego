package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sqlitego/engine"
	"strings"
)

const cliName string = "sqlitego"

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func printUnknown(text string) {
	fmt.Println("'", text, "'", ": command not found")
}

func printHelp() {
	fmt.Printf(
		"Welcome to %v! These are the available commands: \n",
		cliName,
	)
	fmt.Println("\n")
	fmt.Println("\t help           - Show available commands")
	fmt.Println("\t clear          - Clear the terminal screen")
	fmt.Println("\t exit    - Closes your connection to <dbname>")
	fmt.Println("\n")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

func handleCmd(command string) {
	commandPrefix := strings.Split(command, " ")[0]
	switch commandPrefix {
	case "":
		return
	case "help":
		printHelp()
	case "clear":
		clearScreen()
	case "select":
		engine.HandleSelectStatement(command)
	case "insert":
		engine.HandleInsertStatement(command)
	case "exit", "quit":
		os.Exit(0)
	default:
		handleInvalidCmd(command)
	}
}

func parseInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func InitRepl() {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		command := parseInput(reader.Text())
		handleCmd(command)
		printPrompt()
	}

}
