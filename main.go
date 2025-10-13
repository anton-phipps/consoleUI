package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"japanton.com/consoleUI/internal/constants"
	"japanton.com/consoleUI/internal/images"
	"japanton.com/consoleUI/internal/interactions"
	"japanton.com/consoleUI/internal/uiElements"
)

func init() {
	fmt.Printf("%sH%sJ", constants.ESC, constants.ESC)
}

func main() {
	argc := os.Args
	var xDim, yDim int
	if len(argc) < 2 {
		fmt.Println("You must have the name of the image file as an argument.")
		os.Exit(1)
	}
	fileName := argc[1]
	if fileName[len(fileName)-4:] != ".png" {
		fmt.Println("The file must be an '.png'.")
		os.Exit(1)
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatal("file does not exist")
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
	fmt.Println("Loading image, press any key to continue...")
	fmt.Println("...")

	images.LoadPng(fileName, xDim, yDim)
	element := uiElements.NewMenu(0, 0, 100, 100, []string{"Anton", "Burnell", "Phipps"})
	fmt.Printf("Enter a key...\n")
	k, err := interactions.ReadKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Key was read. %s\n", k)
	for {
		fmt.Println("Enter a character...")
		str, err := interactions.ReadKey()
		ch := byte(str[0])
		if err != nil {
			log.Fatal(err)
		}
		if ch == 'q' {
			break
		}
		switch int(ch) {
		case constants.UPARROW:
			if element.Current <= 0 {
				element.Current = 0
			} else {
				element.Current--
			}
		case constants.DOWNARROW:
			if element.Current >= len(element.Items) {
				element.Current = len(element.Items)
			} else {
				element.Current++
			}
		}
		fmt.Printf("%+v\n", element)
		fmt.Printf("value of key: %d\n", ch)
	}
}
