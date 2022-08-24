package main

import "fmt"

func main() {
	/*
		DENGAN MENULISKAN TIPE DATA
	*/
	var firstName string = "John"
	var lastName string
	lastName = "Wick"
	/*
		TANPA MENULISKAN TIPE DATA
	*/
	// var firstName = "John"
	// lastName := "Wick"
	fmt.Printf("Halo, %s %s!\n", firstName, lastName)
}
