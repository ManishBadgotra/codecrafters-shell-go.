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
			cmd := cmds[1]
			cmds[1] = RemoveSingleQuote(cmd)
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
			cmd := cmds[1]
			cmds[1] = RemoveSingleQuote(cmd)
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

func RemoveSingleQuote(input string) string {
	// s := strings.Trim(input, "\r\n")
	// var tokenOfS []string
	// for {
	// 	start := strings.Index(s, "'")
	// 	if start == -1 {
	// 		tokenOfS = append(tokenOfS, strings.Split()...)
	// 		break
	// 	}
	// 	tokenOfS = append(tokenOfS, strings.Fields(s[:start])...)
	// 	s = s[start+1:]
	// 	end := strings.Index(s, "'")
	// 	token := s[:end]
	// 	tokenOfS = append(tokenOfS, token)
	// 	s = s[end+1:]
	// }

	input = strings.Split(input, "'")[1]
	input = strings.Split(input, "'")[0]

	return input
}
