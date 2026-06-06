package main

import "fmt"

func printMenu() {
	fmt.Printf(`
 |+++++++++++++++++++++++++++++++++++++++++++++++|
 |++++++++++++++Maintenance Tracker++++++++++++++|
 | 1. Verfügbare Fahrzeuge
 | 2. Fahrzeug entfernen
 | 3. Fahrzeug Status
 | 4. Inspektion eintragen
 | 5. Arbeit eintragen
 | 6. Fahrzeug hinzufügen
 |
 | q. Menü verlassen
`)
}

func printContinue() {
	fmt.Println("\nDrücke Enter um fortzufahren")
}

func printAddVehicle() {
	fmt.Println(`Um ein Fahrzeug aufzunehmen sind folgende Angaben benötigt:
Marke, Modell, Modellcode, Jahrgang, Farbe, Typ, Kraftstoff, Laufleistung, Kennzeichen`)
}

func wrongSintax() {
	fmt.Printf(`This input isn't allowed`)
}

func printGarage(garage []Vehicle) {
	for i, vehicle := range garage {
		// year is formatted to only show the month and the year instead of the full date
		yearFormatted := vehicle.Year.Format("Jan 2006")

		fmt.Printf(`%d.
		Marke: %s,
		Modell: %s,
		Farbe: %s,
		Typ: %s,
		Kraftstoff: %s,
		Jahrgang: %s,
		Km-Stand: %d,
		Kennzeichen: %s`,
			i+1, vehicle.Brand, vehicle.Model, vehicle.Color, vehicle.Type, vehicle.Fuel, yearFormatted, vehicle.Mileage, vehicle.LicensePlate)
	}
}

func printSelectToErase() {
	fmt.Println(`Bitte ID eingeben, um entsprechendes Fahrzeug zu löschen`)
}
func printEraseSuccessfull(id string) {
	fmt.Printf("Fahrzeug %s wurde gelöscht", id)
}

func printSelectToStatus() {
	fmt.Println(`Bitte ID eingeben, um entsprechenden Fahrzeugstatus zu sehen`)
}

func printVehicleStatus(vehicle Vehicle) {
	fmt.Printf(`
Marke: %s
Modell: %s
Farbe: %s
Typ: %s
Fuel: %s
Jahrgang: %s
Km-Stand: %d
Kennzeichen: %s
`, vehicle.Brand, vehicle.Model, vehicle.Color, vehicle.Type, vehicle.Fuel, vehicle.Year.Format("Jan 2006"), vehicle.Mileage, vehicle.LicensePlate)
}

func printAddWork() {
	fmt.Println(`Um eine Arbeit einzutragen sind folgende Angaben benötigt:
ID, Kategorie, Beschreibung, Teile, Kosten für Teile, Arbeitskosten`)
}

func printGoodbye() {
	fmt.Printf("Tschüss!")
}

func printAddInspection() {
	fmt.Println(`Um eine Inspektion einzutragen sind folgende Angaben benötigt:
ID, Km-Stand, Kosten, Kategorie, Beschreibung, Teile, Kosten für Teile`)
}
