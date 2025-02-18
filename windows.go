//go:build windows
// +build windows

package main

/*
#include <conio.h>

char getChar()
{
	return getch();
}
*/
import "C"
import (
	"fmt"
	"os"
	"syscall"
)

func init() {
	stdout := syscall.Handle(os.Stdout.Fd())
	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004
	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
	fmt.Printf("%sH%sJ", ESC, ESC)
}

func getChar() byte {
	return byte(C.getChar())
}
