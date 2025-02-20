//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	fmt.Printf("%sH%sJ", ESC, ESC)
}

func getChar() byte {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	return b[0]
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
