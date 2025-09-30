//go:build !windows && !darwin
// +build !windows,!darwin

package osSpecific

/*
#include <termios.h>
#include <unistd.h>
#include <stdio.h>
char GetChar()
{
	struct termios oldt, newt;
	char ch;
	tcgetattr(STDIN_FILENO, &oldt);
	newt = oldt;
	newt.c_lflag&= ~(ICANON | ECHO);
	tcsetattr(STDIN_FILENO, TCSANOW, &newt);
	ch = getchar();
	tcsetattr(STDIN_FILENO, TCSANOW, &oldt);
	return ch;
}
*/
import "C"
import (
	"fmt"
	"os"
	"os/exec"
)

func GetChar() byte {
	return byte(C.GetChar())
}

func GetConsoleSize() (rows, cols int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		return 0, 0, err
	}
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return rows, cols, nil
}
