package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

func parseCommand(input string) {
	switch {
	case input == "q":
		os.Exit(0)
	case input == "1":
		fmt.Println(vehicles)
		printContinue()
		clearTerminal()
		printMenu()
	case input == "2":
		clearTerminal()

	}
}

func executeCommand() {
	command := askForCommand() // take a command
	parseCommand(command)      // parse
}

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
