package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		commands, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stdout, "error reading input:"+err.Error())
			log.Println(err, "found while getting command from user.")
			// break
			os.Exit(1)
		}
		commands = strings.ToLower(commands[:len(commands)-1])

		args := strings.Split(commands, " ")
		command := args[0]

		switch command {
		case "exit":
			handleExit(args[1:])
		case "echo":
			executeEcho(args[1:])
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

func handleExit(args []string) {
	exitCode, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stdout, "error while processing error code."+err.Error())
	}
	os.Exit(exitCode)
	// fmt.Println()
}

func executeEcho(args []string) {
	fmt.Println(strings.Trim(fmt.Sprint(args), "[]"))
}
