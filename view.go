package main

import "fmt"

func printMenu() {
	fmt.Println(`
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
