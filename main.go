package main

import (
	"fmt"
	"os"
)

func main() {
	argc := os.Args
	if len(argc) != 2 {
		fmt.Println("You must have the name of the file as an argument.")
		os.Exit(1)
	}
	fileName := argc[1]
	fmt.Println(fileName)
	if fileName[len(fileName)-4:] != ".png" {
		fmt.Println("The file must be an '.png'.")
		os.Exit(1)
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}

	rows, cols, err := getConsoleSize()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Rows:", rows, "Columns:", cols)
	loadPng(fileName)
}
