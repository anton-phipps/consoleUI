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
	"os/exec"
	"strings"
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

func getConsoleSize() (rows, cols int, err error) {
	cmd := exec.Command("mode", "con", "/status")
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Columns") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				colsStr := parts[1]
				fmt.Sscanf(colsStr, "%d", &cols)
			}
		} else if strings.Contains(line, "Lines") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				rowsStr := parts[1]
				fmt.Sscanf(rowsStr, "%d", &rows)
			}
		}
	}

	return rows, cols, nil
}
