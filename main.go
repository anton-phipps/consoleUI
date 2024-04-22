package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argc := os.Args
	var xDim, yDim int
	if len(argc) < 2 {
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

	if len(argc) >= 5 {
		switch argc[2] {
		case "-dim":
			if x, err := strconv.Atoi(argc[3]); err != nil {
				fmt.Println("Invalid X Dimension.")
			} else {
				xDim = x
			}
			if y, err := strconv.Atoi(argc[4]); err != nil {
				fmt.Println("Invalid Y Dimension.")
			} else {
				yDim = y
			}
		default:
			xDim = 0
			yDim = 0
		}
	}
	loadPng(fileName, xDim, yDim)
}
