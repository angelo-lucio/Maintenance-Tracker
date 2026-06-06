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
		printGoodbye()
		os.Exit(0)
	case input == "1":
		clearTerminal()
		allVehicles := garage()
		printGarage(allVehicles)
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case input == "2":
		clearTerminal()
		printGarage(garage())
		printSelectToErase()
		askForCommand()
		eraseVehicle()
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case input == "3":
		clearTerminal()
		printGarage(garage())
		printSelectToStatus()
		response := askForCommand()
		index, err := strconv.Atoi(response)
		if err != nil {
			wrongSintax()
		}
		if index < 1 || index > len(vehicles) {
			wrongSintax()
		}
		printVehicleStatus(vehicles[index-1])
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case input == "4":
		clearTerminal()
		printGarage(garage())
		printAddInspection()
		response := askForCommand()
		addWork(response, garage())
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
	case input == "5":
		clearTerminal()
		printGarage(garage())
		printAddWork()
		response := askForCommand()
		addWork(response, garage())
		printContinue()
		askForCommand()
		clearTerminal()
		printMenu()
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

func eraseVehicle() {
	response := askForCommand()
	index, err := strconv.Atoi(response)
	if err != nil {
		wrongSintax()
	}
	if index < 1 || index > len(vehicles) {
		wrongSintax()
	}
	vehicles = append(vehicles[:index-1], vehicles[index:]...)
	clearTerminal()
	printMenu()
}

func addWork(response string, garage []Vehicle) {
	clean := strings.ReplaceAll(response, ", ", ",")
	parts := strings.Split(clean, ",")
	// index is the first part of the input, which is used to identify the vehicle to which the work should be added
	index, err := strconv.Atoi(parts[0])
	if err != nil || index < 1 || index > len(garage) {
		wrongSintax()
		return
	}

	// mileage parsed from the input, which is used to identify the mileage at which the work was done
	mileage, err := strconv.Atoi(parts[4])
	if err != nil {
		wrongSintax()
		return
	}

	// costs parse
	costs, err := strconv.ParseFloat(parts[5], 64)
	if err != nil {
		wrongSintax()
		return
	}

	// partCosts parse
	partCosts, err := strconv.ParseFloat(parts[6], 64)
	if err != nil {
		wrongSintax()
		return
	}

	work := Maintenance{
		Mileage: mileage,
		Date:    time.Now(),
		Costs:   costs,
		Work: WorksDetail{
			Category:    parts[1],
			Description: parts[2],
			Parts:       parts[3],
			PartCosts:   partCosts,
		},
	}
	// add the work to the corresponding vehicle in the garage
	garage[index-1].Maintenance = append(garage[index-1].Maintenance, work)
	// update the service interval of the vehicle to the current date
	garage[index-1].ServiceInt = time.Now()
}
