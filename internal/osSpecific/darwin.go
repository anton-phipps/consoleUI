//go:build darwin
// +build darwin

package osSpecific

// To get the current version of getch() to work when compiling on macOS, I had to run the following command:
//  sudo mv /usr/local/include/signal.h /usr/local/include/signal.h.bak

/*
#include <stdio.h>
#include <termios.h>
#include <unistd.h>

int getch() {
    struct termios oldt, newt;
    int ch;
    tcgetattr(STDIN_FILENO, &oldt);
    newt = oldt;
    newt.c_lflag &= ~(ICANON | ECHO);
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
	return byte(C.getch())

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
