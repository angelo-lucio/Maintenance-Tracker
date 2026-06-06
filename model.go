package main

import (
	"time"
)

type Vehicle struct {
	Brand, Model, MotorCode, Color, Type, Fuel, LicensePlate string
	Mileage, Power                                           int
	ServiceInt, Year                                         time.Time
	// Maintenance is a slice of Maintenance structs, which contains all the maintenance work that has been done on the vehicle
	Maintenance []Maintenance
}

var vehicles []Vehicle

func addVehicle(vehicle Vehicle) {
	vehicles = append(vehicles, vehicle)
}

func garage() []Vehicle { return vehicles }

type Maintenance struct {
	Mileage int
	Date    time.Time
	Costs   float64
	Work    WorksDetail
}

type WorksDetail struct {
	Category, Description, Parts string
	PartCosts                    float64
}
