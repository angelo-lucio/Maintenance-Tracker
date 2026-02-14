package main

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
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
		vehicle := createVehicle(response)
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

func createVehicle(response string) Vehicle {
	clean := strings.ReplaceAll(response, ", ", ",")
	parts := strings.Split(clean, ",")
	YearLayout := "1.01.1900"
	Year, err := time.Parse(YearLayout, parts[6])
	if err != nil {
		wrongSintax()
	}
	mileage, err := strconv.Atoi(parts[7])
	if err != nil {
		wrongSintax()
	}

	return Vehicle{
		Brand:     parts[0],
		Model:     parts[1],
		MotorCode: parts[2],
		Color:     parts[3],
		Type:      parts[4],
		Fuel:      parts[5],
		Year:      Year,
		Mileage:   mileage,
	}

}
