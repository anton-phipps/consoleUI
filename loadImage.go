package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const ESC = "\u001b["

func loadPng(fileName string, xDim, yDim int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	displayImage(img, xDim, yDim)
}

func displayCell(background, foreground int) {
	fmt.Printf("%s48;5;%dm%s38;5;%dm%c", ESC, background, ESC, foreground, '\u2584')
	fmt.Printf("%s0m", ESC)
}

func convert256(r, g, b int) int {
	return 16 + (min(5, int(r/51)) * 36) + (min(5, int(g/51)) * 6) + min(5, int(b/51))
}

func getRepresentation(img image.Image) (int, int, int) {
	var r, g, b, count uint32
	for row := range img.Bounds().Max.Y - 1 - img.Bounds().Min.Y {
		for col := range img.Bounds().Max.X - 1 - img.Bounds().Min.X {
			temp_r, temp_g, temp_b, _ := img.At(img.Bounds().Min.X+col, img.Bounds().Min.Y+row).RGBA()
			r += temp_r
			g += temp_g
			b += temp_b
			count += 1
		}
	}
	return int(r / count), int(g / count), int(b / count)
}

func scale32Bit(val int) int {
	return int(float32(val) / 256.0)
}

func displayImage(img image.Image, xDim, yDim int) {
	xMin, yMin := img.Bounds().Min.X, img.Bounds().Min.Y
	xMax, yMax := img.Bounds().Max.X-1, img.Bounds().Max.Y-1
	if xDim == 0 || yDim == 0 {
		rows, cols, err := getConsoleSize()
		if err != nil {
			log.Fatal("Could not get Console Size.", err.Error())
			os.Exit(1)
		}
		if xDim == 0 && yDim == 0 {
			if xMax-xMin <= cols && yMax-yMin <= rows*2 {
				xDim = 1
				yDim = 1
			} else if int(float32(yMax-yMin)/float32(rows*2)) > int(float32(xMax-xMin)/float32(cols)) {
				xDim = int(float32(yMax-yMin)/float32(rows*2)) + 1
				yDim = int(float32(yMax-yMin)/float32(rows*2)) + 1
			} else {
				xDim = int(float32(xMax-xMin)/float32(cols)) + 1
				yDim = int(float32(xMax-xMin)/float32(cols)) + 1
			}
		} else {
			if xDim == 0 {
				xDim = yDim
			}
			if yDim == 0 {
				yDim = xDim
			}
		}
	}
	for r := range int((yMax - yMin) / (yDim * 2)) {
		for c := range int((xMax - xMin) / xDim) {
			sectionA := image.NewRGBA(image.Rect(0, 0, xDim, yDim))
			sectionB := image.NewRGBA(image.Rect(0, 0, xDim, yDim))
			draw.Draw(sectionA, sectionA.Bounds(), img, image.Point{xMin + c*xDim, yMin + r*2*yDim}, draw.Src)
			draw.Draw(sectionB, sectionB.Bounds(), img, image.Point{xMin + c*xDim, yMin + r*2*yDim + yDim}, draw.Src)
			rA, gA, bA := getRepresentation(sectionA)
			rB, gB, bB := getRepresentation(sectionB)
			displayCell(convert256(scale32Bit(rA), scale32Bit(gA), scale32Bit(bA)), convert256(scale32Bit(rB), scale32Bit(gB), scale32Bit(bB)))
		}
		fmt.Println()
	}
}
