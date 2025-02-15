package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		clear_in := strings.TrimRight(in, "\n")
		cmds := strings.Split(clear_in, " ")
		switch cmds[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			// cmdAltered := parseArgs(cmds[1])
			cmdsAltered := RemoveSingleQuote(cmds[1:])
			cmds[1] = cmdsAltered

			fmt.Println(strings.Join(cmds[1:], " "))
		case "type":
			switch cmds[1] {
			case "exit", "echo", "type", "pwd", "cd", "cat":
				fmt.Printf("%s is a shell builtin\n", cmds[1])
			default:
				paths := strings.Split(os.Getenv("PATH"), ":")
				isFound := false
				for _, path := range paths {
					fullPath := filepath.Join(path, cmds[1])
					if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
						fmt.Printf("%s is %s\n", cmds[1], fullPath)
						isFound = true
						break
					}
				}
				if !isFound {
					fmt.Printf("%s: not found\n", cmds[1])
				}
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("found error in printing directory:" + err.Error())
			}
			fmt.Println(dir)
		case "cd":
			if cmds[1] == "~" {
				if err := os.Chdir(os.Getenv("HOME")); err != nil {
					fmt.Fprintln(os.Stdout, "cd: HOME: No such file or directory")
				}
			} else if err := os.Chdir(cmds[1]); err != nil {
				fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", cmds[1])
			}
		case "cat":
			command := exec.Command(cmds[0], cmds[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			if err := command.Run(); err != nil {
				fmt.Printf("%s: command not found\n", cmds[0])
			}
		default:
			command := exec.Command(cmds[0], cmds[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			if err := command.Run(); err != nil {
				fmt.Printf("%s: command not found\n", cmds[0])
			}
		}
	}
}

func RemoveSingleQuote(args []string) string {
	for i, s := range args {
		args[i] = strings.ReplaceAll(s, "'", "")
	}

	s := RemoveExtraSpace(args[0])

	return s
}

func ParseArgs(args string) string {
	// get rid of all extra spaces
	args = strings.TrimSpace(args)
	inSingleQuotes := false
	var parsedArgs []string
	var argsBuilder strings.Builder
	for i := 0; i < len(args); i++ {
		c := args[i]
		if c == '\'' {
			inSingleQuotes = !inSingleQuotes
			continue
		}
		if c == ' ' && !inSingleQuotes {
			if argsBuilder.Len() > 0 {
				parsedArgs = append(parsedArgs, argsBuilder.String())
				argsBuilder.Reset()
			}
			continue
		}
		argsBuilder.WriteByte(c)
	}
	//
	if argsBuilder.Len() > 0 {
		parsedArgs = append(parsedArgs, argsBuilder.String())
		argsBuilder.Reset()
	}

	s := fmt.Sprintf("%v", strings.Join(parsedArgs, ""))
	return s
}

func RemoveExtraSpace(s string) string {
	var result string
	var isFirstSpaceAfterLetter = true
	var space = " "
	for _, l := range s {
		a := string(l)
		if a == space && isFirstSpaceAfterLetter {
			// s[i] = l
			isFirstSpaceAfterLetter = false
			// fmt.Printf("'%v' --- first space\n", a)
			continue
		} else if a != space {
			isFirstSpaceAfterLetter = true
			// fmt.Printf("'%v' --- letter\n", a)
		}
		result = fmt.Sprintf("%v", strings.TrimSpace(a))
	}
	return result
}
