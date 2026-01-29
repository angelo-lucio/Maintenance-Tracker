package main

func main() {
	printMenu() //1. show menu
	for {       // 2. infinity loop
		executeCommand() // 3. wait for input
	}
}
