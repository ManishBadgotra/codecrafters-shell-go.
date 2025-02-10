package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	builtin := [...]string{"exit", "echo", "type"}

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
		case "type":
			typeCheck(builtin, args[1:])
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

func typeCheck(builtin [3]string, args []string) {

	isAvailable := false

	for _, val := range builtin {
		if val == args[0] {
			isAvailable = true
			fmt.Printf("%v is a shell builtin", args[0])
			break
		}
	}

	if !isAvailable {
		fmt.Print("there is something unsual in this command. \nPlease check and try again.\n")
	}
}

func handleExit(args []string) {
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stdout, "error while processing error code."+err.Error())
	}
	os.Exit(exitCode)
	// fmt.Println()
}

func executeEcho(args []string) {
	fmt.Println(strings.Trim(fmt.Sprint(args), "[]"))
}
