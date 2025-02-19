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
