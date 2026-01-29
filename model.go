package main

type Vehicle struct {
	Brand, Model, Code, Color, Type string
	Year, Mileage                   int
	OnlyForRacing                   bool
}

var vehicles []Vehicle

func addVehicle(vehicle Vehicle) {
	vehicles = append(vehicles, vehicle)
}

func garage() []Vehicle {
	return vehicles
}
