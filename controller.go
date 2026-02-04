package main

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
		clearTerminal()
		garage()
		printContinue()
		executeCommand()
		clearTerminal()
		printMenu()
	case input == "2":
		clearTerminal()
	case input == "6":
		clearTerminal()
		printAddVehicle()
		response := askForCommand()
		vehicle := createVehichle(response)
		addVehicle(vehicle)
		clearTerminal()
		printMenu()
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

func createVehichle(response string) Vehicle {
	clean := strings.ReplaceAll(response, ", ", ",")
	parts := strings.Split(clean, ",")
	Year, _ := strconv.Atoi(parts[6])
	mileage, _ := strconv.Atoi(parts[7])

	return Vehicle{
		Brand:         parts[0],
		Model:         parts[1],
		Code:          parts[2],
		Color:         parts[3],
		Type:          parts[4],
		Fuel:          parts[5],
		Year:          Year,
		Mileage:       mileage,
		OnlyForRacing: bool(parts[8] == "1"),
	}

}
