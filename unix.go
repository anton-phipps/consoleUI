//go:build !windows && !darwin
// +build !windows,!darwin

package main

/*
#include <termios.h>
#include <unistd.h>
#include <stdio.h>
char getChar()
{
	struct termios oldt, newt;
	char ch;
	tcgetattr(STDIN_FILENO, &oldt);
	newt = oldt;
	newt.c_lflag&= ~(ICANON | ECHO);
	tcsetattr(STDIN_FILENO, TCSANOW, &newt);
	ch = getchar();
	tcsetattr(STDIN_FILENO, TCSANOW, &oldt);
}
*/
import "C"
import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	fmt.Printf("%sH%sJ", ESC, ESC)
}

func getChar() byte {
	return byte(C.getChar())
}

func getConsoleSize() (rows, cols int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		return 0, 0, err
	}
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return rows, cols, nil
}
