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
		in = strings.TrimRight(in, "\n")
		spaceIncludedArray := strings.Split(in, " ")
		var cmds []string
		for _, cmd := range spaceIncludedArray {
			if cmd != "" {
				cmds = append(cmds, cmd)
			}
		}
		cmds = RemoveSingleQuote(cmds)
		switch cmds[0] {
		case "exit":
			os.Exit(0)
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
		case "echo":
			echo := strings.Join(cmds[1:], " ")
			fmt.Fprintf(os.Stdout, "%v\n", echo)
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
	s := RemoveExtraSpace(args)
	for i, s := range args {
		args[i] = strings.ReplaceAll(s, "'", "")
	}

	// return strings.Join(s, " ")
	return s
}

func RemoveExtraSpace(args []string) []string {
	var isFirstSpaceAfterLetter = true
	var space = " "
	for i, arg := range args {
		var ans string
		for _, l := range arg {
			a := string(l)
			if a == space && isFirstSpaceAfterLetter {
				// s[i] = l
				isFirstSpaceAfterLetter = false
				continue
			} else if a != space {
				isFirstSpaceAfterLetter = true
			}
			ans = fmt.Sprintf("%s%s", ans, a)
		}
		args[i] = ans
	}
	return args
}
