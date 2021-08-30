package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const cliName string = "sqlitego"

func printPrompt() {
    fmt.Print(cliName, "> ")
}
 
func printUnknown(text string) {
    fmt.Println("'",text, "'",": command not found")
}

func displayHelp() {
    fmt.Printf(
        "Welcome to %v! These are the available commands: \n",
        cliName,
    )
    fmt.Println("help           - Show available commands")
    fmt.Println("clear          - Clear the terminal screen")
    fmt.Println("exit | quit    - Closes your connection to <dbname>")
}
 
func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func handleInvalidCmd(text string) {
    defer printUnknown(text)
}

func handleCmd(text string) {
    handleInvalidCmd(text)
}

func parseInput(text string) string {
    output := strings.TrimSpace(text)
    output = strings.ToLower(output)
    return output
}

func main() {
    
    commands := map[string]interface{}{
        "help":  displayHelp,
        "clear": clearScreen,
    }
    
    reader := bufio.NewScanner(os.Stdin)
    printPrompt()
    for reader.Scan() {
        text := parseInput(reader.Text())
        if command, exists := commands[strings.Split(text, " ")[0]]; exists {
            command.(func())()
        } else if strings.EqualFold("exit", text) {
			return
        } else if strings.EqualFold("quit", text) {
			return
        } else {
            handleCmd(text)
        }
        printPrompt()
    }
    
	fmt.Println()
}