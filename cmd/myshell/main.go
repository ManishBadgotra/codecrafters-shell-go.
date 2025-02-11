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
			cmds = RemoveSingleQuote(cmds)
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
			cmds = RemoveSingleQuote(cmds)
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

func RemoveSingleQuote(args []string) []string {
	for i, val := range args {
		fmt.Println("---------------- got " + val)
		if i == 0 {
			continue
		}
		val = strings.ReplaceAll(val, `'`, ``)
		fmt.Println("---------------- changed to this " + val)
		args[i] = val
	}

	return args
}
