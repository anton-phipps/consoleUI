//go:build winodws
// +build winodws

package main

import (
	"fmt"
	"os"
	"syscall"
)

func init() {
	fmt.Println("Windows Version")
	stdout := syscall.Handle(os.Stdout.Fd())
	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004
	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
}
