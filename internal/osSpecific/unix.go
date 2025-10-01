//go:build darwin || linux || freebsd
// +build darwin linux freebsd

package osSpecific

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func GetChar() (byte, error) {
	// Save terminal state
	oldState, err := unix.IoctlGetTermios(0, unix.TIOCGETA)
	if err != nil {
		return 0, err
	}
	newState := *oldState
	newState.Lflag &^= unix.ICANON | unix.ECHO
	newState.Cc[unix.VMIN] = 1
	newState.Cc[unix.VTIME] = 0

	if err := unix.IoctlSetTermios(0, unix.TIOCSETA, &newState); err != nil {
		return 0, err
	}
	defer unix.IoctlSetTermios(0, unix.TIOCSETA, oldState)

	// Read one byte
	buf := make([]byte, 1)
	_, err = unix.Read(0, buf)
	return buf[0], err
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
