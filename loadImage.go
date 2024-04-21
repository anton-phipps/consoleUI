package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

const ESC = "\u001b["

func loadPng(fileName string) {
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
	displayImage(img)
}

func displayCell(background, foreground int) {
	fmt.Printf("%s48;5;%dm%s38;5;%dm%c%s0m", ESC, background, ESC, foreground, '\u2584', ESC)
}

func convert256(r, g, b int) int {
	return 16 + (min(5, int(r/51)) * 36) + (min(5, int(g/51)) * 6) + min(5, int(b/51))
}

func displayImage(img image.Image) {
	xMin, yMin := img.Bounds().Min.X, img.Bounds().Min.Y
	xMax, yMax := img.Bounds().Max.X-1, img.Bounds().Max.Y-1
	for y := yMin; y < yMax; y = y + 4 {
		for x := xMin; x < xMax; x = x + 2 {
			r0, g0, b0, _ := img.At(x, y).RGBA()
			r1, g1, b1, _ := img.At(x+1, y).RGBA()
			r2, g2, b2, _ := img.At(x, y+1).RGBA()
			r3, g3, b3, _ := img.At(x+1, y+1).RGBA()
			rA := int(float32(r0+r1+r2+r3) / (4.0 * 256.0))
			gA := int(float32(g0+g1+g2+g3) / (4.0 * 256.0))
			bA := int(float32(b0+b1+b2+b3) / (4.0 * 256.0))
			r0, g0, b0, _ = img.At(x, y+2).RGBA()
			r1, g1, b1, _ = img.At(x+1, y+2).RGBA()
			r2, g2, b2, _ = img.At(x, y+3).RGBA()
			r3, g3, b3, _ = img.At(x+1, y+3).RGBA()
			rB := int(float32(r0+r1+r2+r3) / (4.0 * 256.0))
			gB := int(float32(g0+g1+g2+g3) / (4.0 * 256.0))
			bB := int(float32(b0+b1+b2+b3) / (4.0 * 256.0))
			displayCell(convert256(rA, gA, bA), convert256(rB, gB, bB))
		}
		fmt.Println()
	}
}
