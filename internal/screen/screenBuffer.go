package screen

import (
	"fmt"
	"image"

	"japanton.com/consoleUI/internal/constants"
)

type ScreenBuffer struct {
	Dimension       ScreenDimension
	BufferDimension ScreenDimension
	CurrentPosition Position
	Buffer          []Cell
	EnforceBorder   bool
	Locked          bool
}

type ScreenDimension struct {
	X int
	Y int
}

type Position struct {
	X int
	Y int
}

type Cell struct {
	BgColor   int
	FgColor   int
	Character rune
}

func DisplayCell(background, foreground int) {
	fmt.Printf("%s48;5;%dm%s38;5;%dm%c", constants.ESC, background, constants.ESC, foreground, constants.LOWERBLOCK)
	fmt.Printf("%s", constants.RESETFORMAT)
}

func CreateBuffer(xdim, ydim, xpos, ypos int) *ScreenBuffer {
	var sb *ScreenBuffer = new(ScreenBuffer)

	return sb
}

func (sc *ScreenBuffer) PrintToScreen() {
	for r := sc.CurrentPosition.Y; r < sc.CurrentPosition.Y+sc.BufferDimension.Y; r++ {
		for c := sc.CurrentPosition.X; c < sc.CurrentPosition.X+sc.BufferDimension.X; c++ {
			// Display the item in:
			_ = sc.Buffer[r*sc.BufferDimension.Y+c]
		}
	}
}

func (sc *ScreenBuffer) ImageToBuffer(img image.Image, width, height int) {

}

func (sc *ScreenBuffer) ClearBuffer() {
	// Fill screen with black background and space as char
}

func (sc *ScreenBuffer) SetBufferBackground(red, green, blue int) {
	// Fill the screen with the color specified
}

func (sc *ScreenBuffer) ResizeUpdate() {
	// When window is resized rewrite the buffer in the window
}

func (sc *ScreenBuffer) ShiftBufferAbsolute(x, y int) {
	// Shift the buffer to have top left set to the x and y passed
}

func (sc *ScreenBuffer) ShiftBufferRelative(x, y int) {
	// adjust the top left of the buffer by the integers specified.
	// negative is up and left, positive down and right
}
